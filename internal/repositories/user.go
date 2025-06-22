package repositories

import (
	"events-system/internal/domain"
	"fmt"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(data domain.UserData) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
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

func (ur *UserRepository) CreateUser(data domain.UserData) (*domain.User, error) {
	user, err := ur.factory.CreateUser(data)

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

// func GetUsers() {}

// func DeleteUser() {}

// func (ur *UserRepository) UpdateUser(id string, data domain.UserData) (*domain.User, error) {

// }
