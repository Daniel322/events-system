package commands

import (
	"context"
	"events-system/internal/components/vo"
	"events-system/internal/domain/account"
	"events-system/internal/domain/user"
	"events-system/pkg/utils"
	"log"
	"os"
)

type ICreateUser struct {
	logger   *log.Logger
	userRepo *user.UserRepo
	accRepo  *account.AccRepo
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

var CreateUser *ICreateUser

func InitCreateUser() {
	var logger = log.New(os.Stdout, "CreateUser ", log.LstdFlags)

	CreateUser = &ICreateUser{
		userRepo: user.Repository,
		accRepo:  account.Repository,
		logger:   logger,
	}
}

func (this ICreateUser) Format(user user.Entity) user.Output {
	return user.ToOutput()
}

func (this ICreateUser) Validate(data CreateUserData) (*CreateUserState, error) {
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

func (this ICreateUser) Run(
	ctx context.Context,
	state CreateUserState,
) (*user.Entity, error) {
	isCurrentTransaction := false
	if ctx.Value("transaction") == nil {
		ctx = this.userRepo.Repository.CreateTransaction(ctx)

		isCurrentTransaction = true
	}

	user := user.New(state.Username)
	acc := account.New(state.AccountValue, state.Type, user.ID)

	err := this.userRepo.Save(ctx, user.ToPlain())

	if err != nil {
		ctx = this.userRepo.Repository.Rollback(ctx)
		return nil, utils.GenerateError("Create user", err.Error())
	}

	err = this.accRepo.Save(ctx, acc.ToPlain())

	if err != nil {
		ctx = this.userRepo.Repository.Rollback(ctx)
		return nil, utils.GenerateError("Create user", err.Error())
	}

	user.AddAccount(acc)

	if isCurrentTransaction {
		ctx = this.userRepo.Repository.Commit(ctx)
	}

	return &user, nil
}
