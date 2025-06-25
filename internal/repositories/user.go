package repositories

import (
	"errors"
	"events-system/internal/domain"
	"events-system/internal/utils"
	"fmt"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(data domain.UserData) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	DeleteUser(id string) (bool, error)
	GetUsers(options map[string]interface{}) (*[]domain.User, error)
	UpdateUser(id string, data domain.UserData) (*domain.User, error)
}

type UserRepository struct {
	Name    string
	db      *gorm.DB
	factory domain.IUserFactory
}

func NewUserRepository(name string, db *gorm.DB, factory domain.IUserFactory) *UserRepository {
	return &UserRepository{
		Name:    name,
		db:      db,
		factory: factory,
	}
}

// TODO: add transaction support in create, update and delete methods

func (ur *UserRepository) CreateUser(data domain.UserData) (*domain.User, error) {
	user, err := ur.factory.Create(data)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// change to value from context
	result := ur.db.Create(user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) GetUserById(id string) (*domain.User, error) {
	user := new(domain.User)

	result := ur.db.First(user, "id =?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) DeleteUser(id string) (bool, error) {
	parsedId, _, err := utils.ParseId(id)

	if err != nil {
		return false, errors.New(err.Error())
	}

	user := domain.User{ID: parsedId}
	result := ur.db.Delete(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return false, result.Error
	}

	return true, nil
}

func (ur *UserRepository) GetUsers(options map[string]interface{}) (*[]domain.User, error) {
	var users *[]domain.User

	result := ur.db.Where(options).Find(&users)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (ur *UserRepository) UpdateUser(id string, data domain.UserData) (*domain.User, error) {
	user, err := ur.GetUserById(id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user, err = ur.factory.Update(user, data)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := ur.db.Save(user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return user, nil
}
