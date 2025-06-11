package usecases

import (
	"errors"
	"events-system/domain"
	"time"

	"github.com/google/uuid"
)

type CreateUserData struct {
	Username string `json:"username"`
}

func CreateUser(data CreateUserData) (*domain.User, error) {
	var id uuid.UUID = uuid.New()

	if len(data.Username) == 0 {
		return nil, errors.New("username cant be empty")
	}

	var user = domain.User{
		ID:        id,
		Username:  data.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}
