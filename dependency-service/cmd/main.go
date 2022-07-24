package main

import (
	"github.com/ancore09/dependency-service/internal/config"
	server2 "github.com/ancore09/dependency-service/internal/server"
)

func main() {
	server := server2.NewGRPCServer()

	config.ReadConfigYML("config.yml")
	cfg := config.GetConfigInstance()

	server.Start(&cfg)
}
