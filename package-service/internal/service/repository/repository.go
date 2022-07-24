package repository

import (
	"context"
	"errors"
	"github.com/ancore09/package-service/internal/service/model"
)

var packs = model.Packages{
	{
		Id:   1,
		Name: "go",
	},
	{
		Id:   2,
		Name: "libbrew",
	},
	{
		Id:   3,
		Name: "libgdx",
	},
}

var id uint64 = 4

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) GetPackageByName(ctx context.Context, name string) (*model.Package, error) {
	pack := packs.FilterByName(name)
	if pack == nil {
		return nil, errors.New("not found")
	}
	return pack, nil
}

func (r *Repository) GetAllPackages(ctx context.Context) ([]*model.Package, error) {
	return packs, nil
}

func (r *Repository) CreatePackage(ctx context.Context, name string) (*model.Package, error) {
	pack := model.Package{
		Id:   id,
		Name: name,
	}

	id += 1

	packs = append(packs, &pack)
	return &pack, nil
}
