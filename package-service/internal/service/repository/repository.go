package repository

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	package_service "github.com/ancore09/package-manager-backend/package-service/pkg/package-service"
	"github.com/ancore09/package-service/internal/service/model"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetPackageByName(ctx context.Context, name string) (*model.Package, error) {
	pack := packs.FilterByName(name)
	if pack == nil {
		return nil, errors.New("not found")
	}
	return pack, nil
}

func (r *Repository) GetAllPackages(ctx context.Context) ([]*model.Package, error) {
	query := sq.Select("*").PlaceholderFormat(sq.Dollar).From("packages").RunWith(r.DB.DB)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res []*model.Package
	err = r.DB.SelectContext(ctx, &res, s, args...)

	return res, err
}

func (r *Repository) CreatePackage(ctx context.Context, name string) (*model.Package, error) {
	query := sq.Insert("packages").Columns("name").PlaceholderFormat(sq.Dollar).Values(name).Suffix("returning id").RunWith(r.DB)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var id uint64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return nil, err
		}

		return &model.Package{
			Id:   id,
			Name: name,
		}, nil
	} else {
		return nil, sql.ErrNoRows
	}
}

func (r *Repository) UpdatePackage(ctx context.Context, pack *package_service.Package) (*model.Package, error) {
	query := sq.Update("packages").Set("name", pack.GetName()).PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": pack.GetId()}).RunWith(r.DB)

	_, err := query.ExecContext(ctx)

	if err != nil {
		return nil, err
	}

	return &model.Package{
		Id:   pack.GetId(),
		Name: pack.GetName(),
	}, nil
}

func (r *Repository) DeletePackage(ctx context.Context, id uint64) error {
	query := sq.Delete("packages").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": id}).RunWith(r.DB)

	_, err := query.ExecContext(ctx)

	if err != nil {
		return err
	}
	return nil
}
