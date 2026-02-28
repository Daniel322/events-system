package tg_commands

import (
	"context"
	"events-system/internal/application/queries"
	"events-system/internal/domain/account"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCmd(ctx context.Context, msg *tgbotapi.MessageConfig, update tgbotapi.Update) {

	accountId := update.Message.From.ID
	t, err := account.NewAccountType("telegram")

	if err != nil {
	}

	state := queries.NewCheckAccountState(string(accountId), t)

	currentAcc, err := queries.CheckAccount.Run(ctx, state)

	fmt.Println(currentAcc)
}
