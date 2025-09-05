package repository

import (
	"events-system/infrastructure/providers/db"
	"events-system/pkg/utils"
	"reflect"
	"strings"
)

type Repository[Entity any] struct {
	*BaseRepository
	Name      string
	TableName string
}

func NewRepository[Entity any](table_name ModelName, base_repository *BaseRepository) *Repository[Entity] {
	return &Repository[Entity]{
		Name:           strings.Title(modelNames[table_name]) + " repository",
		TableName:      modelNames[table_name],
		BaseRepository: base_repository,
	}
}

func (repo *Repository[Entity]) checkTransactionExistance(transaction db.DatabaseInstance) db.DatabaseInstance {
	var instanceForExec db.DatabaseInstance
	if reflect.ValueOf(transaction).Elem().IsValid() {
		instanceForExec = transaction
	} else {
		instanceForExec = repo.BaseRepository.DB.Instance
	}

	return instanceForExec
}

func (repo *Repository[Entity]) Save(entity Entity, transaction db.DatabaseInstance) (*Entity, error) {
	dbTransactionForQueryExec := repo.checkTransactionExistance(transaction)
	savedEntity := dbTransactionForQueryExec.Table(repo.TableName).Save(entity)

	if savedEntity.Error != nil {
		return nil, utils.GenerateError(repo.Name, savedEntity.Error.Error())
	}

	return &entity, nil
}

func (repo *Repository[Entity]) Destroy(id string, transaction db.DatabaseInstance) (bool, error) {
	dbTransactionForQueryExec := repo.checkTransactionExistance(transaction)
	result := dbTransactionForQueryExec.Table(repo.TableName).Delete(id)

	if result.Error != nil {
		return false, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return true, nil
}

func (repo *Repository[Entity]) Find(options map[string]interface{}) (*[]Entity, error) {
	var entities *[]Entity
	result := repo.BaseRepository.DB.Instance.Table(repo.TableName).Find(entities, options)

	if result.Error != nil {
		return nil, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return entities, nil
}
