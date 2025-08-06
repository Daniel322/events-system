package repositories

import (
	"events-system/infrastructure/providers/db"
	"events-system/internal/interfaces"
	"events-system/internal/utils"
	"log"
	"reflect"
)

type Repository[Entity any, CreateData any, UpdateData any] struct {
	Name    string
	db      *db.Database
	factory interfaces.Factory[Entity, CreateData, UpdateData]
}

func NewRepository[
	Entity any,
	CreateData any,
	UpdateData any,
](
	name string,
	db *db.Database,
	factory interfaces.Factory[Entity, CreateData, UpdateData],
) *Repository[Entity, CreateData, UpdateData] {
	return &Repository[Entity, CreateData, UpdateData]{
		Name:    name,
		db:      db,
		factory: factory,
	}
}

func (repo *Repository[Entity, CreateData, UpdateData]) Create(
	data CreateData,
	transaction db.DatabaseInstance,
) (*Entity, error) {
	factoryResult, err := repo.factory.Create(data)

	if err != nil {
		return nil, utils.GenerateError(repo.Name, err.Error())
	}

	var instanceForExec = repo.checkTransactionExistance(transaction)

	result := instanceForExec.Create(factoryResult)

	if result.Error != nil {
		return nil, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return factoryResult, nil
}

func (repo *Repository[Entity, CreateData, UpdateData]) GetById(id string) (*Entity, error) {
	factoryInstance := new(Entity)

	result := repo.db.Instance.First(factoryInstance, "id =?", id)

	if result.Error != nil {
		return nil, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return factoryInstance, nil
}

func (repo *Repository[Entity, CreateData, UpdateData]) Delete(id string, transaction db.DatabaseInstance) (bool, error) {
	parsedId, _, err := utils.ParseId(id)

	if err != nil {
		return false, utils.GenerateError(repo.Name, err.Error())
	}

	var instanceForExec = repo.checkTransactionExistance(transaction)

	entity := new(Entity)
	result := instanceForExec.Where("id = ?", parsedId).Delete(&entity)

	if result.Error != nil {
		return false, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return true, nil
}

func (repo *Repository[Entity, CreateData, UpdateData]) GetList(options map[string]interface{}) (*[]Entity, error) {
	var entities *[]Entity

	result := repo.db.Instance.Where(options).Find(&entities)

	if result.Error != nil {
		return nil, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return entities, nil
}

func (repo *Repository[Entity, CreateData, UpdateData]) Update(
	id string,
	data UpdateData,
	transaction db.DatabaseInstance,
) (*Entity, error) {
	entity, err := repo.GetById(id)

	if err != nil {
		return nil, utils.GenerateError(repo.Name, err.Error())
	}

	entity, err = repo.factory.Update(entity, data)

	if err != nil {
		return nil, utils.GenerateError(repo.Name, err.Error())
	}

	var instanceForExec = repo.checkTransactionExistance(transaction)

	result := instanceForExec.Save(entity)

	if result.Error != nil {
		return nil, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return entity, nil
}

func (repo *Repository[Entity, CreateData, UpdateData]) checkTransactionExistance(transaction db.DatabaseInstance) db.DatabaseInstance {
	var instanceForExec db.DatabaseInstance

	if reflect.ValueOf(transaction).Elem().IsValid() {
		log.SetPrefix("INFO ")
		log.Println(repo.Name + ": " + "transaction coming")
		instanceForExec = transaction
	} else {
		log.SetPrefix("INFO ")
		log.Println(repo.Name + ": " + "transaction not exist")
		instanceForExec = repo.db.Instance
	}

	return instanceForExec
}
