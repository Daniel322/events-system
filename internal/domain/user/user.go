package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  UsernameVO
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(val string) (*User, error) {
	username, err := NewUsername(val)

	if err != nil {
		return nil, err
	}

	return &User{
		Username:  *username,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
