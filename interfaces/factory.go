package interfaces

import entities "events-system/internal/entity"

type UserFactory interface {
	Create(username string) (*entities.User, error)
	Update(user *entities.User, username string) (*entities.User, error)
}
