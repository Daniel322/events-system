package usecases

import (
	"events-system/domain"

	"github.com/google/uuid"
)

type UserService interface {
	GetUser(id uuid.UUID) (*domain.User, error)
	CreateUser(data UserData) (*domain.User, error)
	UpdateUser(user *domain.User, data UserData) (*domain.User, error)
	DeleteUser(id uuid.UUID) (*domain.User, error)
}
