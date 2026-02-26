package account

import (
	"context"
	"events-system/interfaces"
	"events-system/internal/components"
	"events-system/pkg/utils"
)

type AccRepo struct {
	components.Factory
}

var Repository *AccRepo

func InitRepo(repo interfaces.Repository) {
	Repository = &AccRepo{
		Factory: *components.NewFactory("Account", repo),
	}
}

func (r AccRepo) Save(ctx context.Context, value interface{}) error {
	ctx = context.WithValue(ctx, "tableName", "accounts")

	return r.Repository.Save(ctx, value)
}

func (r AccRepo) FindOne(ctx context.Context, options map[string]interface{}) (*Plain, error) {
	accs := new([]Plain)
	ctx = context.WithValue(ctx, "tableName", "accounts")
	ctx = context.WithValue(ctx, "ptr", accs)
	err := r.Repository.Find(ctx, options)

	if err != nil {
		return nil, utils.GenerateError("AccRepo FindOne", err.Error())
	}
	if len(*accs) == 0 {
		return nil, nil
	}

	acc := (*accs)[0]

	return &acc, nil
}
