package repositories

import "events-system/internal/domain"

type IUserRepository interface {
	CreateUser(data domain.UserData) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	DeleteUser(id string) (bool, error)
	GetUsers(options map[string]interface{}) (*[]domain.User, error)
	UpdateUser(id string, data domain.UserData) (*domain.User, error)
}

type Repository[Entity any, CreateData any, UpdateData any] interface {
	Create(data CreateData) (*Entity, error)
	Update(id string, data UpdateData) (*Entity, error)
	GetById(id string) (*Entity, error)
	GetList(options map[string]interface{}) (*[]Entity, error)
	Delete(id string) (bool, error)
}
