package service

import (
	"context"
	"github.com/ancore09/package-service/internal/service/model"
)

type PackageService struct {
	repository PackageRepositoryInterface
}

type PackageRepositoryInterface interface {
	GetAllPackages(ctx context.Context) ([]*model.Package, error)
	GetPackageByName(ctx context.Context, name string) (*model.Package, error)
	CreatePackage(ctx context.Context, name string) (*model.Package, error)
}

func NewPackageService(repo PackageRepositoryInterface) PackageService {
	return PackageService{repository: repo}
}

func (s *PackageService) GetPackageByName(ctx context.Context, name string) (*model.Package, error) {
	pack, err := s.repository.GetPackageByName(ctx, name)

	if err != nil {
		return nil, err
	}

	return pack, nil
}

func (s *PackageService) CreatePackage(ctx context.Context, name string) (*model.Package, error) {
	pack, err := s.repository.CreatePackage(ctx, name)

	if err != nil {
		return nil, err
	}

	return pack, nil
}

func (s *PackageService) GetPackages(ctx context.Context) ([]*model.Package, error) {
	packs, err := s.repository.GetAllPackages(ctx)

	if err != nil {
		return nil, err
	}

	return packs, nil
}
