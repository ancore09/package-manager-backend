package app

import (
	"context"
	package_service "github.com/ancore09/package-manager-backend/package-service/pkg/package-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) DeletePackage(ctx context.Context, req *package_service.DeletePackageRequest) (*package_service.DeletePackageResponse, error) {
	err := i.service.DeletePackage(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return &package_service.DeletePackageResponse{Completed: true}, nil
}
