package services

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/repositories"
	"fmt"
)

type IUserService interface {
	CreateUser(data CreateUserData) (*domain.User, error)
	GetUser(id string) (*domain.User, error)
}

type CreateUserData struct {
	Username  string
	Type      string
	AccountId string
}

type UserService struct {
	Name           string
	DB             *db.Database
	userRepository repositories.IUserRepository
	accRepository  repositories.IAccountRepository
}

func NewUserService(db *db.Database, userRepository repositories.IUserRepository, accRepository repositories.IAccountRepository) *UserService {
	return &UserService{
		Name:           "UserService",
		DB:             db,
		userRepository: userRepository,
		accRepository:  accRepository,
	}
}

func (us UserService) CreateUser(data CreateUserData) (*domain.User, error) {
	transaction := us.DB.CreateTransaction()

	fmt.Println(transaction)

	user, err := us.userRepository.Create(domain.UserData{Username: data.Username}, transaction)

	// var account

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	transaction.Commit()
	return user, nil
}

func (us UserService) GetUser(id string) (*domain.User, error) {
	user, err := us.userRepository.GetById(id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return user, nil
}
