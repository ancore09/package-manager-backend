package app

import (
	"github.com/ancore09/dependency-service/internal/service"
	dependency_service "github.com/ancore09/dependency-service/pkg/dependency-service"
)

type Implementation struct {
	dependency_service.UnimplementedPackageServiceServer
	service service.DependencyService
}

func NewDependencyService(service service.DependencyService) *Implementation {
	return &Implementation{service: service}
}
