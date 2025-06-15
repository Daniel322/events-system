package services

import (
	"errors"
	"events-system/internal/domain"
	"time"

	"github.com/google/uuid"
)

type UserData struct {
	Username string
}

type UserService struct {
	Name string
}

type IUserService interface {
	CreateUser(data UserData) (*domain.User, error)
	UpdateUser(user *domain.User, data UserData) (*domain.User, error)
}

func NewUserService(name string) *UserService {
	return &UserService{
		Name: name,
	}
}

func (us *UserService) CreateUser(data UserData) (*domain.User, error) {
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

func (us *UserService) UpdateUser(user *domain.User, data UserData) (*domain.User, error) {
	if len(data.Username) == 0 {
		return nil, errors.New("username cant be empty")
	}

	user.Username = data.Username
	user.UpdatedAt = time.Now()

	return user, nil
}
