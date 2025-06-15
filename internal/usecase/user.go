package usecases

import (
	"events-system/internal/domain"
	"events-system/internal/services"
	"fmt"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserUseCase struct {
	db      *gorm.DB
	service services.IUserService
}

func NewUserUseCase(db *gorm.DB, service services.IUserService) *UserUseCase {
	return &UserUseCase{
		db:      db,
		service: service,
	}
}

func (us UserUseCase) CreateUser(data services.UserData) (*domain.User, error) {
	user, err := us.service.CreateUser(data)

	fmt.Println(user)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// change to value from context
	result := us.db.Create(user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return user, nil
}

func (us UserUseCase) GetUser(id string) (*domain.User, error) {
	var user = new(domain.User)
	return user, nil
}
