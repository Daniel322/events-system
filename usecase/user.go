package usecases

import (
	"errors"
	"events-system/domain"
	"time"

	"github.com/google/uuid"
)

type UserData struct {
	Username string `json:"username"`
}

func CreateUser(data UserData) (*domain.User, error) {
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

func UpdateUser(user *domain.User, data UserData) (*domain.User, error) {
	if len(data.Username) == 0 {
		return nil, errors.New("username cant be empty")
	}

	user.Username = data.Username
	user.UpdatedAt = time.Now()

	return user, nil
}
