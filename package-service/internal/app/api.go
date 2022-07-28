package app

import (
	package_service "github.com/ancore09/package-manager-backend/package-service/pkg/package-service"
	"github.com/ancore09/package-service/internal/service"
)

type Implementation struct {
	package_service.UnimplementedPackageServiceServer
	service service.PackageService
}

func NewPackageService(service service.PackageService) *Implementation {
	return &Implementation{service: service}
}
