package tg_commands

import (
	"context"
	"fmt"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UploadCmd(
	ctx context.Context,
	msg *tgbotapi.MessageConfig,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
) error {
	relativePath := "infrastructure/static/example.csv"
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	docMsg := tgbotapi.NewDocument(update.Message.From.ID, tgbotapi.FilePath(absPath))
	bot.Send(docMsg)
	(*msg).Text = "You can to upload many events in one file\nnow supported only csv format\nuse the example from the file to avoid errors."
	return nil
}
