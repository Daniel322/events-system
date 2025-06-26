package repositories

import "events-system/internal/domain"

type Repository[Entity any, CreateData any, UpdateData any] interface {
	Create(data CreateData) (*Entity, error)
	Update(id string, data UpdateData) (*Entity, error)
	GetById(id string) (*Entity, error)
	GetList(options map[string]interface{}) (*[]Entity, error)
	Delete(id string) (bool, error)
}

type IUserRepository interface {
	Repository[domain.User, domain.UserData, domain.UserData]
}
