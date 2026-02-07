package user

import (
	"context"
	"events-system/interfaces"
	"events-system/internal/components"
)

type UserRepo struct {
	components.Factory
}

func NewUsersRepo(repo interfaces.RepositoryV2) *UserRepo {
	return &UserRepo{
		Factory: *components.NewFactory("User", repo),
	}
}

func (r UserRepo) Save(ctx context.Context, value interface{}) error {
	ctx = context.WithValue(ctx, "tableName", "users")

	return r.Repository.Save(ctx, value)
}
