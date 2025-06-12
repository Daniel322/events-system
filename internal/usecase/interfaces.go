package usecases

import (
	"context"
	"events-system/internal/domain"
	"events-system/internal/services"

	"github.com/google/uuid"
)

type IUserService interface {
	GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error)
	CreateUser(ctx context.Context, data services.UserData) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User, data services.UserData) (*domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
