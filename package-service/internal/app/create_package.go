package app

import (
	"context"
	package_service "github.com/ancore09/package-manager-backend/package-service/pkg/package-service"
	"github.com/ancore09/package-service/internal/service/model"
)

func (i *Implementation) CreatePackage(ctx context.Context, req *package_service.CreatePackageRequest) (*package_service.CreatePackageResponse, error) {
	pack, err := i.service.CreatePackage(ctx, req.GetName())

	if err != nil {
		return nil, err
	}
	return CreatePackageResponse(pack), nil
}

func CreatePackageResponse(pack *model.Package) *package_service.CreatePackageResponse {
	return &package_service.CreatePackageResponse{
		Package: makeServicePackage(pack),
	}
}
