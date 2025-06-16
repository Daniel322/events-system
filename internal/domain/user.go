package domain

import (
	"errors"
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

type IUserFactory interface {
	CreateUser(data UserData) (*User, error)
	UpdateUser(user *User, data UserData) (*User, error)
}

func NewUserFactory(name string) *UserFactory {
	return &UserFactory{
		Name: name,
	}
}

func (us *UserFactory) CreateUser(data UserData) (*User, error) {
	var id uuid.UUID = uuid.New()

	if len(data.Username) == 0 {
		return nil, errors.New("username cant be empty")
	}

	var user = User{
		ID:        id,
		Username:  data.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (us *UserFactory) UpdateUser(user *User, data UserData) (*User, error) {
	if len(data.Username) == 0 {
		return nil, errors.New("username cant be empty")
	}

	user.Username = data.Username
	user.UpdatedAt = time.Now()

	return user, nil
}
