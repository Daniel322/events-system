package usecases

import (
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
)

func (usecase *InternalUseCases) CreateUser(data dto.CreateUserInput) (*dto.OutputUser, error) {
	transaction := usecase.BaseRepository.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	user, err := usecase.UserService.Create(data.Username, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError("CreateUser", err.Error())
	}

	acc, err := usecase.AccountService.Create(dto.CreateAccountData{
		UserId:    user.ID,
		AccountId: data.AccountId,
		Type:      data.Type,
	}, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError("CreateUser", err.Error())
	}

	var accs []entities.Account
	accs = append(accs, *acc)

	if trRes := transaction.Commit(); trRes.Error != nil {
		return nil, utils.GenerateError("CreateUser", trRes.Error.Error())
	}

	return &dto.OutputUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  accs,
	}, nil
}
