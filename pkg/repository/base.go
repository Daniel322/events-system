package repository

import (
	"events-system/infrastructure/providers/db"
	dependency_container "events-system/pkg/di"
)

type BaseRepository struct {
	DB *db.Database
}

func NewBaseRepository(DB *db.Database) {
	dependency_container.Container.Add(
		"baseRepository",
		BaseRepository{
			DB: DB,
		},
	)
}

func (base *BaseRepository) CreateTransaction() db.DatabaseInstance {
	ctx := base.DB.Instance.Begin()

	return ctx
}
