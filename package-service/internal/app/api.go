package app

import (
	"github.com/ancore09/package-service/internal/service"
	package_service "github.com/ancore09/package-service/pkg/package-service"
)

type Implementation struct {
	package_service.UnimplementedPackageServiceServer
	service service.PackageService
}

func NewPackageService(service service.PackageService) *Implementation {
	return &Implementation{service: service}
}
