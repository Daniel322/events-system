package tg_handlers

import (
	"context"
	"events-system/internal/application/queries"
	"events-system/internal/domain/account"
	"events-system/pkg/utils"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CheckAccHandler(ctx context.Context, msg *tgbotapi.MessageConfig, update tgbotapi.Update) (*account.Plain, error) {
	currentAccountId := update.Message.From.ID
	strAccId := strconv.Itoa(int(currentAccountId))

	t, err := account.NewAccountType("telegram")

	checkAccState, err := queries.NewCheckAccountState(strAccId, t)

	if err != nil {
		return nil, utils.GenerateError("CheckAccHandler", err.Error())
	}

	currentAcc, err := queries.CheckAccount.Run(ctx, *checkAccState)

	if err != nil {
		return nil, utils.GenerateError("CheckAccHandler", err.Error())
	}

	if currentAcc == nil {
		msg.Text = "need to create account, use /start command"
		return nil, utils.GenerateError("CheckAccHandler", "need to create account, use /start command")
	}

	return currentAcc, nil
}
