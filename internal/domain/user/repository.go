package user

import (
	"context"
	"events-system/interfaces"
	"events-system/internal/components"
	"events-system/pkg/utils"
	"fmt"
)

type UserRepo struct {
	components.Factory
}

func NewUsersRepo(repo interfaces.Repository) *UserRepo {
	return &UserRepo{
		Factory: *components.NewFactory("User", repo),
	}
}

func (r UserRepo) Save(ctx context.Context, value interface{}) error {
	ctx = context.WithValue(ctx, "tableName", "users")

	return r.Repository.Save(ctx, value)
}

func (r UserRepo) FindOne(ctx context.Context, options map[string]interface{}) (*Plain, error) {
	users := new([]Plain)
	ctx = context.WithValue(ctx, "tableName", "users")
	ctx = context.WithValue(ctx, "ptr", users)
	err := r.Repository.Find(ctx, options)

	fmt.Println((*users)[0])

	if err != nil {
		return nil, utils.GenerateError("UserRepo", err.Error())
	}

	user := (*users)[0]

	return &user, nil
}
