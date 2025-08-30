package repository

import (
	"events-system/infrastructure/providers/db"
	"events-system/pkg/utils"
	"strings"
)

type Repository[Entity any] struct {
	Name      string
	TableName string
	DB        db.Database
}

func NewRepository[Entity any](DB db.Database, tableName ModelName) *Repository[Entity] {
	return &Repository[Entity]{
		Name:      strings.Title(modelNames[tableName]) + " repository",
		TableName: modelNames[tableName],
		DB:        DB,
	}
}

func (repo *Repository[Entity]) Save(entity Entity, transaction db.DatabaseInstance) (*Entity, error) {
	// TODO: check incoming transaction instance
	savedEntity := repo.DB.Instance.Table(repo.TableName).Save(entity)

	if savedEntity.Error != nil {

	}

	return &entity, nil
}

func (repo *Repository[Entity]) Destroy(id string) (bool, error) {
	result := repo.DB.Instance.Table(repo.TableName).Delete(id)

	if result.Error != nil {
		return false, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return true, nil
}
