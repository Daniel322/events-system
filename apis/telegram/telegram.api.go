package telegram_api

import (
	"context"
	account_module "events-system/modules/account"
	"events-system/modules/db"
	event_module "events-system/modules/event"
	user_module "events-system/modules/user"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI
var err error
var eventsChatSlice []int64

func BootstrapBot() {
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			fmt.Println("length of chats slice", len(eventsChatSlice))
			for i, v := range eventsChatSlice {
				if v == update.Message.Chat.ID {
					fmt.Println("match case", update.Message.Text)
					currentUserId, err := account_module.GetUserIdByAccountId(strconv.Itoa(int(update.Message.From.ID)))
					if err != nil {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
						msg.Text = "user not found, account not created. use /start before work for create account"
						bot.Send(msg)
						continue
					}
					msgSlice := strings.Split(update.Message.Text, `/`)
					event_module.CreateEvent(event_module.CreateEventData{
						Info: msgSlice[1],
						// TODO: fix next error: can't scan into dest[3] (col: date): cannot scan timestamp (OID 1114) in binary format into *string
						Date:      msgSlice[0],
						UserId:    currentUserId,
						Providers: []string{"telegram"},
					}, context.Background())
					eventsChatSlice = append(eventsChatSlice[:i], eventsChatSlice...)
					fmt.Println("length of chats slice after delete", len(eventsChatSlice))
				}
			}
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			fmt.Println(update.Message.From.ID)
			fmt.Println(update.Message.From.UserName)
			var currentAccCount int
			currentAccCount, err = account_module.GetAccountByAccountId(strconv.Itoa(int(update.Message.From.ID)))
			if err != nil {
				msg.Text = "Something went wrong"
				break
			}
			if currentAccCount > 0 {
				msg.Text = "You already create account"
				break
			}
			currentContext := context.Background()
			transaction, _ := db.Connection.Begin(currentContext)
			var user, err = user_module.CreateUser(
				user_module.CreateUserData{
					Username: update.Message.From.UserName,
				},
				currentContext,
			)
			if err != nil {
				msg.Text = "Something went wrong"
				transaction.Rollback(currentContext)
				break
			}
			_, err = account_module.CreateAccount(
				account_module.AccountData{
					UserId:    user.Id,
					AccountId: strconv.Itoa(int(update.Message.From.ID)),
					Type:      "telegram",
				},
				currentContext,
			)
			if err != nil {
				msg.Text = "Something went wrong"
				transaction.Rollback(currentContext)
				break
			}
			transaction.Commit(currentContext)
			msg.Text = "Start, account created"
		case "add":
			eventsChatSlice = append(eventsChatSlice, update.Message.Chat.ID)
			msg.Text = "send information about your event in next format: YYYY.MM.DD / event-info"
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
