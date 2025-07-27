package services

import (
	"events-system/internal/domain"
	"events-system/internal/interfaces"
	"events-system/internal/providers/db"
	"events-system/internal/structs"
	"events-system/internal/utils"
)

type UserService struct {
	Name           string
	DB             *db.Database
	userRepository interfaces.IRepository[domain.User, domain.UserData, domain.UserData]
	accRepository  interfaces.IRepository[domain.Account, domain.CreateAccountData, domain.UpdateAccountData]
}

func NewUserService(
	db *db.Database,
	userRepository interfaces.IRepository[domain.User, domain.UserData, domain.UserData],
	accRepository interfaces.IRepository[domain.Account, domain.CreateAccountData, domain.UpdateAccountData],
) *UserService {
	return &UserService{
		Name:           "UserService",
		DB:             db,
		userRepository: userRepository,
		accRepository:  accRepository,
	}
}

func (us UserService) CreateUser(data structs.CreateUserData) (*structs.User, error) {
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

	var accs []domain.Account
	accs = append(accs, *acc)

	if trRes := transaction.Commit(); trRes.Error != nil {
		return nil, utils.GenerateError(us.Name, trRes.Error.Error())
	}

	return &structs.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  accs,
	}, nil
}

func (us UserService) GetUser(id string) (*structs.User, error) {
	user, err := us.userRepository.GetById(id)

	if err != nil {
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	options := make(map[string]interface{})

	options["user_id"] = user.ID

	accs, err := us.accRepository.GetList(options)

	if err != nil {
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	return &structs.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  *accs,
	}, nil
}
