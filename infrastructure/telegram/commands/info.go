package tg_commands

import (
	"context"
	"events-system/internal/application/queries"
	"events-system/internal/domain/account"
	"events-system/pkg/utils"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InfoCmd(
	ctx context.Context,
	msg *tgbotapi.MessageConfig,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
) error {
	accountId := update.Message.From.ID
	strAccId := strconv.Itoa(int(accountId))
	t, err := account.NewAccountType("telegram")

	if err != nil {
		return utils.GenerateError("InfoCmd", err.Error())
	}

	checkAccState := queries.NewCheckAccountState(strAccId, t)

	currentAcc, err := queries.CheckAccount.Run(ctx, checkAccState)

	if err != nil {
		return utils.GenerateError("InfoCmd", err.Error())
	}

	currentUser, err := queries.GetUser.Run(ctx, currentAcc.UserId)

	msg.Text = "Username: " + currentUser.Username + "\n"

	currentEvents, err := queries.EventsList.Run(ctx, currentUser.ID)

	msg.Text += "Events:\n"

	for _, event := range *currentEvents {
		msg.Text += event.Info + " in " + event.Date.Format("2006-01-02") + "\n"
	}

	return nil
}
