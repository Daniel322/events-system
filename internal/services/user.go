package services

import (
	"events-system/internal/domain"
	"events-system/internal/repositories"
	"fmt"
)

type IUserService interface {
	CreateUser(data domain.UserData) (*domain.User, error)
	GetUser(id string) (*domain.User, error)
}

type UserService struct {
	Name           string
	userRepository repositories.IUserRepository
}

func NewUserService(name string, repository repositories.IUserRepository) *UserService {
	return &UserService{
		Name:           name,
		userRepository: repository,
	}
}

func (us UserService) CreateUser(data domain.UserData) (*domain.User, error) {
	// TODO: add logic for create account also
	user, err := us.userRepository.Create(data)

	fmt.Println(user)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

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
