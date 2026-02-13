package account

import (
	"context"
	"events-system/interfaces"
	"events-system/internal/components"
)

type AccRepo struct {
	components.Factory
}

func NewAccRepo(repo interfaces.RepositoryV2) *AccRepo {
	return &AccRepo{
		Factory: *components.NewFactory("Account", repo),
	}
}

func (r AccRepo) Save(ctx context.Context, value interface{}) error {
	ctx = context.WithValue(ctx, "tableName", "accounts")

	return r.Repository.Save(ctx, value)
}

func (r AccRepo) FindOne(ctx context.Context, options map[string]interface{}) (*Entity, error) {
	res, err := r.Repository.Find(ctx, options)

	cur := (*res)[0]

	return cur.(*Entity), err
}
