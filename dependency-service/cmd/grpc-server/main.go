package main

import (
	"context"
	"github.com/ancore09/dependency-service/internal/config"
	server2 "github.com/ancore09/dependency-service/internal/server"
	package_service "github.com/ancore09/package-manager-backend/package-service/pkg/package-service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	server := server2.NewGRPCServer()

	config.ReadConfigYML("config.yml")
	cfg := config.GetConfigInstance()

	packageServiceConn, err := grpc.DialContext(
		context.Background(),
		"localhost:6002",
		grpc.WithInsecure())

	if err != nil {
		log.Error().Err(err).Msg("failed to create client")
	}

	_ = package_service.NewPackageServiceClient(packageServiceConn)

	server.Start(&cfg)
}
