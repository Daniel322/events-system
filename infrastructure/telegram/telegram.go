package telegram

import (
	"context"
	"events-system/infrastructure/config"
	pg_db "events-system/infrastructure/db/adapters/postgres"
	parsers "events-system/infrastructure/parser"
	"events-system/internal/application/commands"
	"fmt"

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

		if !update.Message.IsCommand() {
			if update.Message.Document != nil {
				reader, err := tg_handlers.FileHandler(ctx, &msg, update, tg.Bot)

				if err != nil {
					tg.Bot.Send(msg)
					continue
				}

				currentAcc, err := tg_handlers.CheckAccHandler(ctx, &msg, update)

				if err != nil {
					tg.Bot.Send(msg)
					continue
				}

				eventsData, err := parsers.ParseCsv(
					ctx,
					reader,
					parsers.ParseOptions{
						AccId:  currentAcc.ID,
						UserId: currentAcc.UserId,
					},
				)

				if err != nil {
					msg.Text = err.Error()
					tg.Bot.Send(msg)
					continue
				}

				// TODO: подумать как можно вынести логику транзакции из этого слоя, чтобы был какой-то хендлер, куда можно передать контекст и колбек, хендлер будет отвечать за ролбеки и комиты
				if ctx.Value("transaction") == nil {
					ctx = pg_db.Adapter.CreateTransaction(ctx)
				}

				for _, eventData := range *eventsData {
					state, err := commands.CreateEvent.Validate(eventData)

					if err != nil {
						ctx = pg_db.Adapter.Rollback(ctx)
						tg.Bot.Send(msg)
						continue
					}

					_, err = commands.CreateEvent.Run(ctx, state)

					if err != nil {
						ctx = pg_db.Adapter.Rollback(ctx)
						tg.Bot.Send(msg)
						continue
					}
				}

				ctx = pg_db.Adapter.Commit(ctx)

				msg.Text = fmt.Sprint(len(*eventsData)) + " events created!"
			}
			if len(update.Message.Text) != 0 {
				tg_handlers.MessageHandler(ctx, &msg, update)
			}
		} else {
			cb, ok := tg_commands.COMMANDS[update.Message.Command()]
			if !ok {
				tg_commands.DefaultCmd(ctx, &msg, update)
			}
			err := cb(ctx, &msg, update, tg.Bot)

			if err != nil {
				msg.Text = err.Error()
				tg.Bot.Send(msg)
				continue
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
