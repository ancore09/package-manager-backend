package app

import (
	dependency_service "github.com/ancore09/dependency-service/pkg/dependency-service"
)

type Implementation struct {
	dependency_service.UnimplementedPackageServiceServer
}

//func NewDependencyService(service service.PackageService) *Implementation {
//	return &Implementation{service: service}
//}
