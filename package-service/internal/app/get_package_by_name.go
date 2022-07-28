package app

import (
	"context"
	package_service "github.com/ancore09/package-manager-backend/package-service/pkg/package-service"
	"github.com/ancore09/package-service/internal/service/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetPackageByName(ctx context.Context, req *package_service.GetPackageByNameRequest) (*package_service.GetPackageByNameResponse, error) {
	pack, err := i.service.GetPackageByName(ctx, req.GetName())
	if err != nil || pack == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return makeGetPackageByNameResponse(pack), nil

}

func makeGetPackageByNameResponse(pack *model.Package) *package_service.GetPackageByNameResponse {
	return &package_service.GetPackageByNameResponse{
		Package: &package_service.Package{Id: pack.Id, Name: pack.Name},
	}
}
