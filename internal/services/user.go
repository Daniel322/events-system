package services

import (
	"events-system/internal/domain"
	"events-system/internal/repositories"
	"fmt"
)

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (us UserService) CreateUser(data domain.UserData) (*domain.User, error) {
	// TODO: add logic for create account also
	user, err := us.userRepository.CreateUser(data)

	fmt.Println(user)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return user, nil
}

func (us UserService) GetUser(id string) (*domain.User, error) {
	var user = new(domain.User)
	return user, nil
}
