package main

import (
	"backend/api"
	db "backend/db/sqlc"
	"backend/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

//const (
//	dbDriver      = "postgres"
//	dbSource      = "postgresql://root:root@localhost:5433/dcoker?sslmode=disable"
//	serverAddress = "0.0.0.0:8080"
//)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("fatal error config file: %w", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot open db connection ", err)
	}

	server, err := api.NewServer(config, *db.New(conn))

	if err != nil {
		log.Fatal("Cannot create server ", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server ", err)
	}
}
