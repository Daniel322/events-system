package repositories

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/utils"
	"log"
	"reflect"
)

type AccountRepository struct {
	Name    string
	db      *db.Database
	factory domain.IAccountFactory
}

func NewAccountRepository(db *db.Database, factory domain.IAccountFactory) *AccountRepository {
	return &AccountRepository{
		Name:    "AccountRepository",
		db:      db,
		factory: factory,
	}
}

func (ar *AccountRepository) Create(data domain.CreateAccountData, transaction db.DatabaseInstance) (*domain.Account, error) {
	_, strV, err := utils.ParseId(data.UserId)

	if err != nil {
		return nil, utils.GenerateError(ar.Name, err.Error())
	}

	acc, err := ar.factory.Create(domain.CreateAccountData{
		UserId:    strV,
		AccountId: data.AccountId,
		Type:      data.Type,
	})

	if err != nil {
		return nil, utils.GenerateError(ar.Name, err.Error())
	}

	var instanceForExec db.DatabaseInstance

	if reflect.ValueOf(transaction).Elem().IsValid() {
		log.SetPrefix("INFO ")
		log.Println(ar.Name + ": " + "transaction coming")
		instanceForExec = transaction
	} else {
		log.SetPrefix("INFO ")
		log.Println(ar.Name + ": " + "transaction not exist")
		instanceForExec = ar.db.Instance
	}

	result := instanceForExec.Create(acc)

	if result.Error != nil {
		return nil, utils.GenerateError(ar.Name, result.Error.Error())
	}

	return acc, nil
}

func (ar *AccountRepository) GetById(id string) (*domain.Account, error) {
	acc := new(domain.Account)

	result := ar.db.Instance.First(acc, "id = ?", id)

	if result.Error != nil {
		return nil, utils.GenerateError(ar.Name, result.Error.Error())
	}

	return acc, nil
}

func (ar *AccountRepository) GetList(options map[string]interface{}) (*[]domain.Account, error) {
	var accs = new([]domain.Account)

	result := ar.db.Instance.Where(options).Find(accs)

	if result.Error != nil {
		return nil, utils.GenerateError(ar.Name, result.Error.Error())
	}

	return accs, nil
}

func (ar *AccountRepository) Update(id string, data domain.UpdateAccountData) (*domain.Account, error) {
	acc, err := ar.GetById(id)

	if err != nil {
		return nil, utils.GenerateError(ar.Name, err.Error())
	}

	acc, err = ar.factory.Update(acc, data)

	if err != nil {
		return nil, utils.GenerateError(ar.Name, err.Error())
	}

	result := ar.db.Instance.Save(acc)

	if result.Error != nil {
		return nil, utils.GenerateError(ar.Name, result.Error.Error())
	}

	return acc, nil
}

func (ar *AccountRepository) Delete(id string) (bool, error) {
	parsedId, _, err := utils.ParseId(id)

	if err != nil {
		return false, utils.GenerateError(ar.Name, err.Error())
	}

	acc := new(domain.Account)
	acc.ID = parsedId
	result := ar.db.Instance.Delete(acc)

	if result.Error != nil {
		return false, utils.GenerateError(ar.Name, result.Error.Error())
	}

	return true, nil
}
