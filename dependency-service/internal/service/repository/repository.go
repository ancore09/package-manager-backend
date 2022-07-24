package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ancore09/dependency-service/internal/service/model"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"log"
)

type Repository struct {
	dgraphClient *dgo.Dgraph
}

func New() *Repository {
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{
		dgraphClient: dgo.NewDgraphClient(api.NewDgraphClient(d)),
	}
}

func (r *Repository) GetDepsForPackage(ctx context.Context, id uint64) (*model.Tree, error) {
	var q string = `
		{
			q(func: eq(id, %d)) @recurse {
				id
				dependOn
			}
		}
	`

	q = fmt.Sprintf(q, id)

	resp, err := r.dgraphClient.NewTxn().Query(ctx, q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp.GetJson()))

	decode := model.Tree{}

	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Printf("%+v\n", decode)

	return &decode, nil
}
