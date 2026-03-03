package telegram

import (
	"context"
	"events-system/infrastructure/config"
	tg_commands "events-system/infrastructure/telegram/commands"
	tg_handlers "events-system/infrastructure/telegram/handlers"
	"events-system/pkg/utils"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBotProvider struct {
	Logger *log.Logger
	Bot    *tgbotapi.BotAPI
}

var Provider *TgBotProvider

func NewTgBotProvider() error {
	token, err := config.Config.TG_TOKEN()

	if err != nil {
		return utils.GenerateError("TgProvider", err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return utils.GenerateError("TgProvider", err.Error())
	}

	var logger = log.New(os.Stdout, "TgProvider"+" ", log.LstdFlags)

	Provider = &TgBotProvider{
		Logger: logger,
		Bot:    bot,
	}

	return nil
}

func (tg *TgBotProvider) Bootstrap() {
	tg.Logger.Printf("Authorized on account %s", tg.Bot.Self.UserName)

	ctx := context.Background()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		tg.Logger.Println(update.Message.From.ID)

		if !update.Message.IsCommand() {
			tg_handlers.MessageHandler(ctx, &msg, update)
		} else {
			switch update.Message.Command() {
			// here will be our cmd handlers
			// TODO: make map with commands, create interface and change switch case for get command from map
			case "help":
				tg_commands.HelpCmd(ctx, &msg)
			case "start":
				err := tg_commands.StartCmd(ctx, &msg, update)

				if err != nil {
					msg.Text = err.Error()
					tg.Bot.Send(msg)
					continue
				}
			case "event":
				tg_commands.EventCmd(ctx, &msg, update)
			default:
				tg_commands.DefaultCmd(ctx, &msg)
			}
		}

		if len(msg.Text) == 0 {
			msg.Text = "Something went wrong"
		}

		if _, err := tg.Bot.Send(msg); err != nil {
			utils.GenerateError("TgProvider", err.Error())
		}
	}
}

func (tg *TgBotProvider) Send(chatId int64, text string) {
	tg.Logger.Println(chatId, text)
	msg := tgbotapi.NewMessage(chatId, text)

	m, err := tg.Bot.Send(msg)

	tg.Logger.Println(err, m)
}

func (tg *TgBotProvider) Close() {
	tg.Bot.StopReceivingUpdates()
	tg.Logger.Println("close tg provider")
}
