package components

import (
	"events-system/interfaces"
	entities "events-system/internal/entity"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Factory
}

func NewUsersFactory(repo interfaces.RepositoryV2) *Users {
	return &Users{
		Factory: *NewFactory("User", repo),
	}
}

func (factory Users) NewUser(username string) *User {
	id := uuid.New()
	created := time.Now()
	factory.Logger.Println("Create user", username, "with id", id)
	return &User{
		User: entities.User{
			ID:        uuid.New(),
			CreatedAt: created,
			UpdatedAt: created,
			Username:  username,
		},
		RepositoryV2: factory.Repository,
	}
}
