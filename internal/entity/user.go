package entities

import (
	dependency_container "events-system/pkg/di"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserFactory struct {
	Name string
}

const USERNAME_CANT_BE_EMPTY_ERR_MSG = "username cant be empty"

func NewUserFactory() *UserFactory {
	factory := &UserFactory{
		Name: "UserFactory",
	}

	dependency_container.Container.Add("userFactory", factory)

	return factory
}

func (us *UserFactory) Update(user *User, username string) (*User, error) {
	err := us.checkUsername(username)

	if err != nil {
		return nil, err
	}

	user.Username = username
	user.UpdatedAt = time.Now()

	return user, nil
}
