package queries

import (
	"context"
	"events-system/internal/domain/account"
	"log"
	"os"
)

type ICheckAccount struct {
	logger  *log.Logger
	accRepo *account.AccRepo
}

type CheckAccountState struct {
	value account.AccountValue
}

func NewCheckAccountState(value string, t account.AccountType) CheckAccountState {
	accValue, err := account.NewAccountValue(value, t)

	if err != nil {
		// TODO: add handle err
	}

	return CheckAccountState{value: accValue}
}

var CheckAccount *ICheckAccount

func InitCheckAccount() {
	var logger = log.New(os.Stdout, "CheckAcc ", log.LstdFlags)
	CheckAccount = &ICheckAccount{
		logger:  logger,
		accRepo: account.Repository,
	}
}

func (this ICheckAccount) Run(
	ctx context.Context,
	state CheckAccountState,
) (*account.Plain, error) {
	options := make(map[string]interface{})
	options["account_id"] = state.value.Val()
	acc, err := this.accRepo.FindOne(ctx, options)

	return acc, err
}
