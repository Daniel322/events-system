package repositories

import (
	"events-system/infrastructure/providers/db"
	"events-system/internal/utils"
	"log"
	"reflect"
)

var connection *db.Database

func Init(conn *db.Database) {
	connection = conn
}

// TODO: make private function for dynamic generate name for logs

func checkTransactionExistance(transaction db.DatabaseInstance, name string) db.DatabaseInstance {
	var instanceForExec db.DatabaseInstance

	if reflect.ValueOf(transaction).Elem().IsValid() {
		log.SetPrefix("INFO ")
		log.Println(name + " repository: " + "transaction coming")
		instanceForExec = transaction
	} else {
		log.SetPrefix("INFO ")
		log.Println(name + "repository: " + "transaction not exist")
		instanceForExec = connection.Instance
	}

	return instanceForExec
}

func Create[Entity any](data Entity, transaction db.DatabaseInstance) (*Entity, error) {
	typeName := reflect.TypeOf(data).String()
	var instanceForExec = checkTransactionExistance(transaction, typeName)

	result := instanceForExec.Create(&data)

	if result.Error != nil {
		return nil, utils.GenerateError(typeName, result.Error.Error())
	}

	return &data, nil
}

func GetById[Entity any](id string) (*Entity, error) {
	entity := new(Entity)
	typeName := reflect.TypeOf(entity).String()

	result := connection.Instance.First(entity, "id =?", id)

	if result.Error != nil {
		return nil, utils.GenerateError(typeName+" repository", result.Error.Error())
	}

	return entity, nil
}

func Delete[Entity any](id string, transaction db.DatabaseInstance) (bool, error) {
	entity := new(Entity)
	typeName := reflect.TypeOf(entity).String()
	parsedId, _, err := utils.ParseId(id)

	if err != nil {
		return false, utils.GenerateError(typeName+" repository", err.Error())
	}

	var instanceForExec = checkTransactionExistance(transaction, typeName)

	result := instanceForExec.Where("id = ?", parsedId).Delete(&entity)

	if result.Error != nil {
		return false, utils.GenerateError(typeName+" repository", result.Error.Error())
	}

	return true, nil
}
