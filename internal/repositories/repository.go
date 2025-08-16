package repositories

import (
	"events-system/infrastructure/providers/db"
	"events-system/internal/utils"
	"log"
	"reflect"
)

type Repository[Entity any] struct {
	Name string
	db   *db.Database
}

func NewRepository[Entity any](
	name string,
	db *db.Database,
) *Repository[Entity] {
	return &Repository[Entity]{
		Name: name,
		db:   db,
	}
}

func (repo *Repository[Entity]) Create(
	data Entity,
	transaction db.DatabaseInstance,
) (*Entity, error) {
	var instanceForExec = repo.checkTransactionExistance(transaction)

	result := instanceForExec.Create(&data)

	if result.Error != nil {
		return nil, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return &data, nil
}

func (repo *Repository[Entity]) GetById(id string) (*Entity, error) {
	entity := new(Entity)

	result := repo.db.Instance.First(entity, "id =?", id)

	if result.Error != nil {
		return nil, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return entity, nil
}

func (repo *Repository[Entity]) Delete(id string, transaction db.DatabaseInstance) (bool, error) {
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

func (repo *Repository[Entity]) GetList(options map[string]interface{}) (*[]Entity, error) {
	var entities *[]Entity

	result := repo.db.Instance.Where(options).Find(&entities)

	if result.Error != nil {
		return nil, utils.GenerateError(repo.Name, result.Error.Error())
	}

	return entities, nil
}

func (repo *Repository[Entity]) Update(
	id string,
	data Entity,
	transaction db.DatabaseInstance,
) (*Entity, error) {
	entity, err := repo.GetById(id)

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

func (repo *Repository[Entity]) checkTransactionExistance(transaction db.DatabaseInstance) db.DatabaseInstance {
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
