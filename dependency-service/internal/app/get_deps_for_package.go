package app

import (
	"context"
	dependency_service "github.com/ancore09/dependency-service/pkg/dependency-service"
)

func (i *Implementation) GetPackageDeps(ctx context.Context, req *dependency_service.GetPackageDepsRequest) (*dependency_service.GetPackageDepsResponse, error) {
	depsForPackage, err := i.service.GetDepsForPackage(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return makeGetPackageDepsResponse(depsForPackage), nil
}

func makeGetPackageDepsResponse(deps []uint64) *dependency_service.GetPackageDepsResponse {
	return &dependency_service.GetPackageDepsResponse{DepsId: deps}
}
