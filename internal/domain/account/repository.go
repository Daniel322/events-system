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

// func (r AccRepo) FindOne(ctx context.Context, options map[string]interface{}) (*Entity, error) {
// 	ctx = context.WithValue(ctx, "tableName", "accounts")
// 	ctx = context.WithValue(ctx, "entityType", reflect.TypeOf(Entity{}))
// 	res, err := r.Repository.Find(ctx, options)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(*res) == 0 {
// 		return nil, nil
// 	}
// 	return (*res)[0].(*Entity), nil
// }
