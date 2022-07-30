package app

import (
	"context"
	package_service "github.com/ancore09/package-manager-backend/package-service/pkg/package-service"
	"github.com/ancore09/package-service/internal/service/model"
)

func (i *Implementation) UpdatePackage(ctx context.Context, req *package_service.UpdatePackageRequest) (*package_service.UpdatePackageResponse, error) {
	pack, err := i.service.UpdatePackage(ctx, req.GetPackage())

	if err != nil {
		return nil, err
	}

	return makeUpdatePackageResponse(pack), nil
}

func makeUpdatePackageResponse(pack *model.Package) *package_service.UpdatePackageResponse {
	return &package_service.UpdatePackageResponse{
		Package: makeServicePackage(pack),
	}
}
