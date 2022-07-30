package service

import (
	"context"
	package_service "github.com/ancore09/package-manager-backend/package-service/pkg/package-service"
	"github.com/ancore09/package-service/internal/service/model"
)

type PackageService struct {
	repository PackageRepositoryInterface
}

type PackageRepositoryInterface interface {
	GetAllPackages(ctx context.Context) ([]*model.Package, error)
	GetPackageByName(ctx context.Context, name string) (*model.Package, error)
	CreatePackage(ctx context.Context, name string) (*model.Package, error)
	UpdatePackage(ctx context.Context, p *package_service.Package) (*model.Package, error)
	DeletePackage(ctx context.Context, id uint64) error
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

func (s *PackageService) UpdatePackage(ctx context.Context, _pack *package_service.Package) (*model.Package, error) {
	pack, err := s.repository.UpdatePackage(ctx, _pack)

	if err != nil {
		return nil, err
	}

	return pack, nil
}

func (s *PackageService) DeletePackage(ctx context.Context, id uint64) error {
	err := s.repository.DeletePackage(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
