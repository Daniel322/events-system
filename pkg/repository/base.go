package repository

import (
	"events-system/infrastructure/providers/db"
)

type BaseRepository struct {
	DB *db.Database
}

func NewBaseRepository(DB *db.Database) *BaseRepository {
	return &BaseRepository{
		DB: DB,
	}
}

func (base *BaseRepository) CreateTransaction() db.DatabaseInstance {
	ctx := base.DB.Instance.Begin()

	return ctx
}
