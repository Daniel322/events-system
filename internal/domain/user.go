package domain

import (
	"events-system/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserData struct {
	Username string
}

type UserFactory struct {
	Name string
}

const USERNAME_CANT_BE_EMPTY_ERR_MSG = "username cant be empty"

func NewUserFactory() *UserFactory {
	return &UserFactory{
		Name: "UserFactory",
	}
}

func (us *UserFactory) Create(data UserData) (*User, error) {
	var id uuid.UUID = uuid.New()

	if len(data.Username) == 0 {
		return nil, utils.GenerateError(us.Name, USERNAME_CANT_BE_EMPTY_ERR_MSG)
	}

	var user = User{
		ID:        id,
		Username:  data.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (us *UserFactory) Update(user *User, data UserData) (*User, error) {
	if len(data.Username) == 0 {
		return nil, utils.GenerateError(us.Name, USERNAME_CANT_BE_EMPTY_ERR_MSG)
	}

	user.Username = data.Username
	user.UpdatedAt = time.Now()

	return user, nil
}
