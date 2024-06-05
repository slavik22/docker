package api

import (
	db "backend/db/sqlc"
	"backend/token"
	"backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config   util.Config
	store    db.Queries
	router   *gin.Engine
	jwtMaker token.JWTMaker
}

func NewServer(config util.Config, store db.Queries) (*Server, error) {

	jwt, err := token.NewJWTMaker(config.SecretKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{
		store:    store,
		jwtMaker: *jwt,
		config:   config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.POST("/users/register", server.registerUser)
	router.POST("/users/login", server.loginUser)

	router.GET("/tutorials", server.listTutorials)
	router.GET("/tutorials/user/:id", server.listUserTutorials)
	router.POST("/tutorials", server.createTutorial)
	router.PUT("/tutorials/:id", server.updateTutorial)
	router.GET("/tutorials/:id", server.getTutorial)
	router.DELETE("/tutorials/:id", server.deleteTutorial)

	router.GET("/ws", handleWebSocket)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
