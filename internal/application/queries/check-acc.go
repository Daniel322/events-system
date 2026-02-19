package queries

import (
	"context"
	"events-system/internal/domain/account"
	"log"
)

type CheckAccount struct {
	Logger  *log.Logger
	AccRepo *account.AccRepo
}

type CheckAccountState struct {
	value account.AccountValue
}

func (this CheckAccount) Run(
	ctx context.Context,
	state CheckAccountState,
) (*account.Plain, error) {
	options := make(map[string]interface{})
	options["account_id"] = state.value.Val()
	acc, err := this.AccRepo.FindOne(ctx, options)

	return acc, err
}
