package components

import (
	"context"
	"events-system/interfaces"
	entities "events-system/internal/entity"
)

type User struct {
	interfaces.RepositoryV2
	entities.User
}

func (user User) get() entities.User {
	return user.User
}

func (comp User) Save(ctx context.Context) error {
	err := comp.RepositoryV2.Save(ctx, comp.User)
	return err
}
func (comp User) Destroy() {}
