package commands

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

type CreateUserData struct {
	Username     string
	Type         string
	AccountValue string
}

type CreateUserState struct {
	Username     vo.NonEmptyString
	Type         account.AccountType
	AccountValue account.AccountValue
}

func NewCreateUser(
	userRepo *user.UserRepo,
	accRepo *account.AccRepo,
) *CreateUser {
	var logger = log.New(os.Stdout, "CreateUser ", log.LstdFlags)

	return &CreateUser{
		UserRepo: userRepo,
		AccRepo:  accRepo,
		Logger:   logger,
	}
}

func (this CreateUser) Format(user user.Entity) user.Output {
	return user.ToOutput()
}

func (this CreateUser) Validate(data CreateUserData) (*CreateUserState, error) {
	state := CreateUserState{}

	username, err := vo.NewNonEmptyString(data.Username)

	if err != nil {
		return nil, utils.GenerateError("CreateUser.Validate", err.Error())
	}

	state.Username = username

	accType, err := account.NewAccountType(data.Type)

	if err != nil {
		return nil, utils.GenerateError("CreateUser.Validate", err.Error())
	}

	state.Type = accType

	accValue, err := account.NewAccountValue(data.AccountValue, accType)

	if err != nil {
		return nil, utils.GenerateError("CreateUser.Validate", err.Error())
	}

	state.AccountValue = accValue

	return &state, nil
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

	user.AddAccount(acc)

	return &user, nil
}
