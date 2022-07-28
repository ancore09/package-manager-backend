package app

import (
	"context"
	package_service "github.com/ancore09/package-manager-backend/package-service/pkg/package-service"
	"github.com/ancore09/package-service/internal/service/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetPackages(ctx context.Context, req *package_service.GetPackagesRequest) (*package_service.GetPackagesResponse, error) {
	packs, err := i.service.GetPackages(ctx)
	if err != nil || packs == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return makeGetPackagesResponse(packs), nil

}

func makeServicePackage(pack *model.Package) *package_service.Package {
	return &package_service.Package{Id: pack.Id, Name: pack.Name}
}

func makeGetPackagesResponse(packs []*model.Package) *package_service.GetPackagesResponse {
	var arr []*package_service.Package

	for _, pack := range packs {
		arr = append(arr, makeServicePackage(pack))
	}

	return &package_service.GetPackagesResponse{
		Package: arr,
	}
}
