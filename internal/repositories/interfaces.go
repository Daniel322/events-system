package repositories

import (
	"events-system/internal/providers/db"
)

type IRepository[Entity any, CreateData any, UpdateData any] interface {
	Create(data CreateData, transaction db.DatabaseInstance) (*Entity, error)
	Update(id string, data UpdateData, transaction db.DatabaseInstance) (*Entity, error)
	GetById(id string) (*Entity, error)
	GetList(options map[string]interface{}) (*[]Entity, error)
	Delete(id string, transaction db.DatabaseInstance) (bool, error)
}
