package repositories

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
)

type Repository[Entity any, CreateData any, UpdateData any] interface {
	Create(data CreateData, transaction db.DatabaseInstance) (*Entity, error)
	// TODO: add transaction support
	Update(id string, data UpdateData) (*Entity, error)
	GetById(id string) (*Entity, error)
	GetList(options map[string]interface{}) (*[]Entity, error)
	// TODO: add transaction support
	Delete(id string) (bool, error)
}

type IUserRepository interface {
	Repository[domain.User, domain.UserData, domain.UserData]
}

type IAccountRepository interface {
	Repository[domain.Account, domain.CreateAccountData, domain.UpdateAccountData]
}
