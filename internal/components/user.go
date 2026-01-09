package components

import (
	"context"
	"events-system/interfaces"
	entities "events-system/internal/entity"
	"time"

	"github.com/google/uuid"
)

type User struct {
	interfaces.RepositoryV2
	entities.User
}

func (user User) get() entities.User {
	return user.User
}

func (user User) Create(ctx *context.Context, username string) entities.User {
	user.User = entities.User{
		Username:  username,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user.Save(ctx, user.User)

	return user.get()
}
func (comp User) Update()  {}
func (comp User) Destroy() {}
