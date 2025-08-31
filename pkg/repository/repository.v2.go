package repository

import (
	"events-system/infrastructure/providers/db"
	dependency_container "events-system/pkg/di"
	"events-system/pkg/utils"
	"reflect"
	"strings"
)

type Repository[Entity any] struct {
	Name       string
	TableName  string
	Connection BaseRepository
}

func NewRepository[Entity any](tableName ModelName) (*Repository[Entity], error) {
	connection, err := dependency_container.Container.Get("baseRepository")

	if err != nil {
		return nil, utils.GenerateError("On init "+modelNames[tableName]+" repository", err.Error())
	}

	return &Repository[Entity]{
		Name:       strings.Title(modelNames[tableName]) + " repository",
		TableName:  modelNames[tableName],
		Connection: (*connection).(BaseRepository),
	}, nil
}

func (repo *Repository[Entity]) checkTransactionExistance(transaction db.DatabaseInstance) db.DatabaseInstance {
	var instanceForExec db.DatabaseInstance
	if reflect.ValueOf(transaction).Elem().IsValid() {
		instanceForExec = transaction
	} else {
		instanceForExec = connection.Instance
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
	result := repo.Connection.DB.Instance.Table(repo.TableName).Find(entities, options)

	if result.Error != nil {
		return nil, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return entities, nil
}
