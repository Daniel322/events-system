package components

import (
	"events-system/interfaces"
	entities "events-system/internal/entity"
	"time"

	"github.com/google/uuid"
)

type UserFactory struct {
	Factory
}

func NewUserFactory(repo interfaces.RepositoryV2) *UserFactory {
	return &UserFactory{
		Factory: *NewFactory("User", repo),
	}
}

func (factory UserFactory) NewUser(username string) *User {
	id := uuid.New()
	created := time.Now()
	factory.Logger.Println("Create user", username, "with id", id)
	return &User{
		User: entities.User{
			Username:  username,
			ID:        uuid.New(),
			CreatedAt: created,
			UpdatedAt: created,
		},
		RepositoryV2: factory.Repository,
	}
}
