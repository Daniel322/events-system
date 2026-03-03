package tg_commands

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DefaultCmd(ctx context.Context, msg *tgbotapi.MessageConfig) {
	msg.Text = "Unknown command"
}
