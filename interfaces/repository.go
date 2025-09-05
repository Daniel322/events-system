package interfaces

import "events-system/infrastructure/providers/db"

type BaseRepository interface {
	CreateTransaction() db.DatabaseInstance
}

// TODO: find sol for change any
type Repository[Entity any] interface {
	BaseRepository
	Save(entity Entity, transaction db.DatabaseInstance) (*Entity, error)
	Destroy(id string, transaction db.DatabaseInstance) (bool, error)
	Find(options map[string]interface{}) (*[]Entity, error)
}
