package services

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/repositories"
	"events-system/internal/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type IUserService interface {
	CreateUser(data CreateUserData) (*User, error)
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

type User struct {
	ID        uuid.UUID
	Username  string
	Accounts  []domain.Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserService(db *db.Database, userRepository repositories.IUserRepository, accRepository repositories.IAccountRepository) *UserService {
	return &UserService{
		Name:           "UserService",
		DB:             db,
		userRepository: userRepository,
		accRepository:  accRepository,
	}
}

func (us UserService) CreateUser(data CreateUserData) (*User, error) {
	transaction := us.DB.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	user, err := us.userRepository.Create(domain.UserData{Username: data.Username}, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	acc, err := us.accRepository.Create(domain.CreateAccountData{
		UserId:    user.ID.String(),
		AccountId: data.AccountId,
		Type:      data.Type,
	}, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	var accs = make([]domain.Account, 1)
	accs = append(accs, *acc)

	if trRes := transaction.Commit(); trRes.Error != nil {
		return nil, utils.GenerateError(us.Name, trRes.Error.Error())
	}

	return &User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  accs,
	}, nil
}

func (us UserService) GetUser(id string) (*domain.User, error) {
	user, err := us.userRepository.GetById(id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return user, nil
}
