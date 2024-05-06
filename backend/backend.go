package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

func checkContainerExists(cli *client.Client, containerName string) (bool, string, error) {
	ctx := context.Background()
	filterArgs := filters.NewArgs()
	filterArgs.Add("name", containerName)
	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true, Filters: filterArgs})
	if err != nil {
		return false, "", fmt.Errorf("error listing containers: %w", err)
	}
	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName {
				return true, container.ID, nil
			}
		}
	}
	return false, "", nil
}

// startDockerDind ensures that a Docker-in-Docker container is running.
func startDockerDind(containerName string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", fmt.Errorf("error creating Docker client: %w", err)
	}
	defer cli.Close()

	// Check if the container already exists
	exists, containerID, err := checkContainerExists(cli, containerName)
	if err != nil {
		return "", err
	}
	if exists {
		fmt.Printf("Container named '%s' already exists with ID %s\n", containerName, containerID)
		return containerID, nil
	}

	ctx := context.Background()
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        "docker:dind",
		Tty:          true,
		ExposedPorts: nat.PortSet{"8181/tcp": struct{}{}},
	}, &container.HostConfig{
		Privileged: true,
		PortBindings: nat.PortMap{
			"8181/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8181",
				},
			},
		},
	}, nil, nil, containerName)
	if err != nil {
		return "", fmt.Errorf("error creating container: %w", err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("error starting container: %w", err)
	}

	return resp.ID, nil
}

// executeDockerCommand runs a command inside the Docker-in-Docker container using the Docker Go SDK.
func executeDockerCommand(containerID string, command []string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", fmt.Errorf("error creating Docker client: %w", err)
	}
	defer cli.Close()

	ctx := context.Background()
	execConfig := types.ExecConfig{
		Cmd:          command,
		AttachStdout: true,
		AttachStderr: true,
	}
	execIDResp, err := cli.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		return "", fmt.Errorf("error creating exec instance: %w", err)
	}

	attachResp, err := cli.ContainerExecAttach(ctx, execIDResp.ID, types.ExecStartCheck{})
	if err != nil {
		return "", fmt.Errorf("error attaching to exec instance: %w", err)
	}
	defer attachResp.Close()

	outputBytes, err := ioutil.ReadAll(attachResp.Reader)
	if err != nil {
		return "", fmt.Errorf("error reading output: %w", err)
	}

	return string(outputBytes), nil
}

// handleWebSocket handles websocket requests from the peer.
func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from browser
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}

		command := strings.Split(strings.ToValidUTF8(string(message), ""), " ")

		containerID, err := startDockerDind("docker-dind")
		if err != nil {
			fmt.Println("Error starting Docker-in-Docker:", err)
			return
		}

		// Execute Docker command
		output, err := executeDockerCommand(containerID, command)
		if err != nil {
			output = fmt.Sprintf("Error executing command: %s", err.Error())
		}

		text := strings.Replace(strings.ToValidUTF8(output, ""), "\n", "\r\n", -1)

		bytes := []byte(text)

		// Send the output back to the browser
		err = conn.WriteMessage(mt, bytes)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
		time.Sleep(time.Second)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ws", handleWebSocket)
	r.Run(":8080")
}
