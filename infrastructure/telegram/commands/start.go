package tg_commands

import (
	"context"
	"events-system/internal/application/commands"
	"events-system/internal/application/queries"
	"events-system/internal/domain/account"
	"events-system/pkg/utils"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCmd(
	ctx context.Context,
	msg *tgbotapi.MessageConfig,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
) error {

	accountId := update.Message.From.ID
	t, err := account.NewAccountType("telegram")

	if err != nil {
		return utils.GenerateError("StartCmd", err.Error())
	}

	strAccId := strconv.Itoa(int(accountId))

	checkAccState := queries.NewCheckAccountState(strAccId, t)

	currentAcc, err := queries.CheckAccount.Run(ctx, checkAccState)

	if err != nil {
		return utils.GenerateError("StartCmd", err.Error())
	}

	if currentAcc != nil {
		currentUser, err := queries.GetUser.Run(ctx, currentAcc.UserId)
		if err != nil {
			return utils.GenerateError("StartCmd", err.Error())
		}
		msg.Text = "account " + currentUser.Username + " already created"
	} else {
		createUserState, err := commands.CreateUser.Validate(commands.CreateUserData{
			Username:     update.Message.From.UserName,
			Type:         "telegram",
			AccountValue: strAccId,
		})

		if err != nil {
			return utils.GenerateError("StartCmd", err.Error())
		}

		newUser, err := commands.CreateUser.Run(ctx, *createUserState)

		msg.Text = "account " + newUser.ToPlain().Username + " created, welcome!"
	}

	return nil
}
