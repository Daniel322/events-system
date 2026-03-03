package interfaces

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command func(ctx context.Context, msg *tgbotapi.MessageConfig, update tgbotapi.Update) error

type Commands map[string]Command
