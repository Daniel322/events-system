package interfaces

import (
	"context"
	"events-system/infrastructure/providers/db"
)

type BaseRepository interface {
	CreateTransaction() db.DatabaseInstance
}

type Repository[Entity any] interface {
	BaseRepository
	Save(entity Entity, transaction db.DatabaseInstance) (*Entity, error)
	Destroy(id string, transaction db.DatabaseInstance) (bool, error)
	Find(options map[string]interface{}) (*[]Entity, error)
}

type RepositoryV2 interface {
	Save(ctx context.Context, value interface{}) error
}
