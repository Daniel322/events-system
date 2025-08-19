package repository

import (
	"events-system/infrastructure/providers/db"
	"events-system/pkg/utils"
	"log"
	"reflect"
)

var connection *db.Database

func Init(conn *db.Database) {
	connection = conn
}

func getGenericName[Entity any](value Entity) string {
	typeName := reflect.TypeOf(value).String()

	return typeName + " repository:"
}

func checkTransactionExistance(transaction db.DatabaseInstance, name string) db.DatabaseInstance {
	var instanceForExec db.DatabaseInstance

	if reflect.ValueOf(transaction).Elem().IsValid() {
		log.SetPrefix("INFO ")
		log.Println(name + "transaction coming")
		instanceForExec = transaction
	} else {
		log.SetPrefix("INFO ")
		log.Println(name + "transaction not exist")
		instanceForExec = connection.Instance
	}

	return instanceForExec
}

func CreateTransaction() db.DatabaseInstance {
	tx := connection.Instance.Begin()

	return tx
}

func Create[Entity any](tableName ModelName, data Entity, transaction db.DatabaseInstance) (*Entity, error) {
	typeName := getGenericName(data)
	var instanceForExec = checkTransactionExistance(transaction, typeName)

	result := instanceForExec.Table(modelNames[tableName]).Create(&data)

	if result.Error != nil {
		return nil, utils.GenerateError(typeName, result.Error.Error())
	}

	return &data, nil
}

func GetById[Entity any](tableName ModelName, id string) (*Entity, error) {
	entity := new(Entity)
	typeName := getGenericName(entity)

	result := connection.Instance.Table(modelNames[tableName]).First(entity, "id =?", id)

	if result.Error != nil {
		return nil, utils.GenerateError(typeName, result.Error.Error())
	}

	return entity, nil
}

func Delete[Entity any](tableName ModelName, id string, transaction db.DatabaseInstance) (bool, error) {
	entity := new(Entity)
	typeName := getGenericName(entity)
	parsedId, _, err := utils.ParseId(id)

	if err != nil {
		return false, utils.GenerateError(typeName, err.Error())
	}

	var instanceForExec = checkTransactionExistance(transaction, typeName)

	result := instanceForExec.Table(modelNames[tableName]).Where("id = ?", parsedId).Delete(&entity)

	if result.Error != nil {
		return false, utils.GenerateError(typeName, result.Error.Error())
	}

	return true, nil
}

func GetList[Entity any](tableName ModelName, options map[string]interface{}) (*[]Entity, error) {
	var entities *[]Entity
	typeName := getGenericName(entities)

	result := connection.Instance.Table(modelNames[tableName]).Where(options).Find(&entities)

	if result.Error != nil {
		return nil, utils.GenerateError(typeName, result.Error.Error())
	}

	return entities, nil
}

func Update[Entity any](
	tableName ModelName,
	id string,
	data Entity,
	transaction db.DatabaseInstance,
) (*Entity, error) {
	entity, err := GetById[Entity](tableName, id)
	typeName := getGenericName(entity)

	if err != nil {
		return nil, utils.GenerateError(typeName, err.Error())
	}

	var instanceForExec = checkTransactionExistance(transaction, typeName)

	result := instanceForExec.Table(modelNames[tableName]).Save(entity)

	if result.Error != nil {
		return nil, utils.GenerateError(typeName, result.Error.Error())
	}

	return entity, nil
}
