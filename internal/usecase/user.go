package usecases

import (
	"events-system/internal/domain"
	db "events-system/internal/providers"
	"events-system/internal/services"
	"log"

	"gorm.io/gorm"
)

type UserUseCase struct {
	Db      *gorm.DB
	Service *services.UserService
}

func NewUserUseCase(db *gorm.DB, service *services.UserService) *UserUseCase {
	return &UserUseCase{
		Db:      db,
		Service: service,
	}
}

func (us UserUseCase) CreateUser(data services.UserData) (*domain.User, error) {
	user, err := us.Service.CreateUser(data)

	if err != nil {
		// log error
	}
	// change to value from context
	result := db.Connection.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return user, nil
}
