package tg_commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HelpCmd(msg *tgbotapi.MessageConfig) {
	// TODO: think about localization
	(*msg).Text = "I understand next commands:\n/start - create acc for you\n/event - create event\n/info - get info about you and your saved events\n/upload - upload file with events"
}
