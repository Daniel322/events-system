package application

import (
	"context"
	"events-system/internal/domain/account"
	"events-system/internal/domain/user"
	"events-system/pkg/utils"
	"events-system/pkg/vo"
	"log"
	"os"
)

type CreateUser struct {
	Logger   *log.Logger
	UserRepo *user.UserRepo
	AccRepo  *account.AccRepo
}

type CreateUserState struct {
	Username     vo.NonEmptyString
	Type         account.AccountType
	AccountValue account.AccountValue
}

func NewCreateUser(userRepo *user.UserRepo, accRepo *account.AccRepo) *CreateUser {
	var logger = log.New(os.Stdout, "CreateUser ", log.LstdFlags)

	return &CreateUser{
		UserRepo: userRepo,
		AccRepo:  accRepo,
		Logger:   logger,
	}
}

func (this CreateUser) Run(
	ctx context.Context,
	state CreateUserState,
) (*user.Entity, error) {
	// TODO: add tranasction

	user := user.New(state.Username)
	acc := account.New(state.AccountValue, state.Type, user.ID)

	err := this.UserRepo.Save(ctx, user.ToPlain())

	if err != nil {
		return nil, utils.GenerateError("Create user", err.Error())
	}

	err = this.AccRepo.Save(ctx, acc.ToPlain())

	if err != nil {
		return nil, utils.GenerateError("Create user", err.Error())
	}

	user.Model.AddAccount(acc.Model)

	return &user, nil
}
