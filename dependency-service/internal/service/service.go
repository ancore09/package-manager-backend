package service

import (
	"context"
	"github.com/ancore09/dependency-service/internal/service/model"
)

type DependencyRepositoryInterface interface {
	GetDepsForPackage(ctx context.Context, id uint64) (*model.Tree, error)
}

type DependencyService struct {
	repository DependencyRepositoryInterface
}

func NewDependencyService(repository DependencyRepositoryInterface) DependencyService {
	return DependencyService{repository: repository}
}

func (s *DependencyService) GetDepsForPackage(ctx context.Context, id uint64) ([]uint64, error) {
	depsTree, err := s.repository.GetDepsForPackage(ctx, id)
	if err != nil {
		return nil, err
	}

	var order []uint64

	dfs(depsTree.Dependencies[0], &order)

	return order, nil
}

func dfs(dependency model.Dependency, order *[]uint64) {
	if len(dependency.DependOn) != 0 {
		for _, dep := range dependency.DependOn {
			dfs(dep, order)
		}
	}
	//fmt.Println(dependency.Id)
	*order = append(*order, dependency.Id)
}
