package tg_commands

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HelpCmd(ctx context.Context, msg *tgbotapi.MessageConfig) {
	// TODO: think about localization
	(*msg).Text = "List of available commands:\n/start - create acc for you\n/event - create event\n/info - get info about you and your saved events\n/upload - upload file with events"
}
