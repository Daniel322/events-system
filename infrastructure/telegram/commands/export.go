package tg_commands

import (
	"context"
	"encoding/csv"
	"events-system/internal/application/queries"
	"events-system/internal/domain/account"
	"events-system/pkg/utils"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ExportCmd(ctx context.Context, msg *tgbotapi.MessageConfig, update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	currentAccountId := update.Message.From.ID
	strAccId := strconv.Itoa(int(currentAccountId))
	t, err := account.NewAccountType("telegram")

	if err != nil {
		return utils.GenerateError("ExportCmd", err.Error())
	}

	checkAccState, err := queries.NewCheckAccountState(strAccId, t)

	if err != nil {
		return utils.GenerateError("ExportCmd", err.Error())
	}

	currentAcc, err := queries.CheckAccount.Run(ctx, *checkAccState)

	if err != nil {
		return utils.GenerateError("ExportCmd", err.Error())
	}

	currentUser, err := queries.GetUser.Run(ctx, currentAcc.UserId)

	if err != nil {
		return utils.GenerateError("ExportCmd", err.Error())
	}

	currentEvents, err := queries.EventsList.Run(ctx, currentUser.ID)

	if err != nil {
		return utils.GenerateError("ExportCmd", err.Error())
	}

	csvData := [][]string{}
	csvData = append(csvData, []string{"Info", "Date"})

	for _, event := range *currentEvents {
		csvData = append(csvData, []string{event.Info, event.Date.Format("2006-01-02")})
	}

	fileName := "events" + "_" + currentUser.Username + "_" + ".csv"

	csvFile, err := os.Create(fileName)

	if err != nil {
		return utils.GenerateError("ExportCmd", err.Error())
	}

	csvWriter := csv.NewWriter(csvFile)

	if err := csvWriter.WriteAll(csvData); err != nil {
		return utils.GenerateError("ExportCmd", err.Error())
	}

	csvWriter.Flush()

	csvFile.Close()

	bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FilePath(fileName)))

	msg.Text = "Events export completed"

	os.Remove(fileName)

	return nil
}
