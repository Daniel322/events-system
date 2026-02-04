package user

import (
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
