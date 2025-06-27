package repositories

import (
	"errors"
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/utils"
	"fmt"
	"log"
	"reflect"
)

type UserRepository struct {
	Name    string
	db      *db.Database
	factory domain.IUserFactory
}

func NewUserRepository(db *db.Database, factory domain.IUserFactory) *UserRepository {
	return &UserRepository{
		Name:    "UserRepository",
		db:      db,
		factory: factory,
	}
}

// TODO: add transaction support in create, update and delete methods

func (ur *UserRepository) Create(data domain.UserData, transaction db.DatabaseInstance) (*domain.User, error) {
	user, err := ur.factory.Create(data)

	if err != nil {
		return nil, utils.GenerateError(ur.Name, err.Error())
	}

	var instanceForExec db.DatabaseInstance

	if reflect.ValueOf(transaction).Elem().IsValid() {
		log.SetPrefix("INFO ")
		log.Println(ur.Name + ": " + "transaction coming")
		instanceForExec = transaction
	} else {
		log.SetPrefix("INFO ")
		log.Println(ur.Name + ": " + "transaction not exist")
		instanceForExec = ur.db.Instance
	}

	result := instanceForExec.Create(user)

	if result.Error != nil {
		return nil, utils.GenerateError(ur.Name, result.Error.Error())
	}

	return user, nil
}

func (ur *UserRepository) GetById(id string) (*domain.User, error) {
	user := new(domain.User)

	result := ur.db.Instance.First(user, "id =?", id)

	if result.Error != nil {
		return nil, utils.GenerateError(ur.Name, result.Error.Error())
	}

	return user, nil
}

func (ur *UserRepository) Delete(id string) (bool, error) {
	parsedId, _, err := utils.ParseId(id)

	if err != nil {
		return false, errors.New(err.Error())
	}

	user := domain.User{ID: parsedId}
	result := ur.db.Instance.Delete(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return false, result.Error
	}

	return true, nil
}

func (ur *UserRepository) GetList(options map[string]interface{}) (*[]domain.User, error) {
	var users *[]domain.User

	result := ur.db.Instance.Where(options).Find(&users)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (ur *UserRepository) Update(id string, data domain.UserData) (*domain.User, error) {
	user, err := ur.GetById(id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user, err = ur.factory.Update(user, data)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := ur.db.Instance.Save(user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return user, nil
}
