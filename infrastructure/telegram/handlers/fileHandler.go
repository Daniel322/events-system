package tg_handlers

import (
	"context"
	"io"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func FileHandler(
	ctx context.Context,
	msg *tgbotapi.MessageConfig,
	update tgbotapi.Update,
	bot *tgbotapi.BotAPI,
) (io.ReadCloser, error) {
	file_id := update.Message.Document.FileID
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: file_id})

	if err != nil {
		msg.Text = err.Error()
		return nil, err
	}

	url := file.Link(bot.Token) // возвращает правильный URL для скачивания

	resp, err := http.Get(url)
	if err != nil {
		msg.Text = err.Error()
		return nil, err
	}

	return resp.Body, nil
}
