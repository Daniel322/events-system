package telegram

import (
	"events-system/internal/services"
	"events-system/internal/utils"
	"log"
	"reflect"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgEvent struct {
	Name string
	Date *time.Time
}

type TgBotProvider struct {
	Name                 string
	Bot                  *tgbotapi.BotAPI
	UserService          services.IUserService
	AccountService       services.IAccountService
	NotCompletedEventMap map[int64]*TgEvent
}

func NewTgBotProvider(
	token string,
	userService services.IUserService,
	accService services.IAccountService,
) (*TgBotProvider, error) {
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return nil, utils.GenerateError("ThBotProvider", err.Error())
	}

	return &TgBotProvider{
		Name:                 "TgBotProvider",
		Bot:                  bot,
		UserService:          userService,
		AccountService:       accService,
		NotCompletedEventMap: make(map[int64]*TgEvent, 10),
	}, nil
}

func (tg *TgBotProvider) Bootstrap() {
	log.SetPrefix("TG_BOT ")
	log.Printf("Authorized on account %s", tg.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			accountId := update.Message.From.ID
			currentAccount, err := tg.AccountService.CheckAccount(accountId)

			if err != nil {
				utils.GenerateError(tg.Name, err.Error())
				break
			}

			if currentAccount == nil {
				strAccId := strconv.Itoa(int(update.Message.From.ID))

				newUser, err := tg.UserService.CreateUser(services.CreateUserData{
					Username:  update.Message.From.UserName,
					AccountId: strAccId,
					Type:      "telegram",
				})

				if err != nil {
					utils.GenerateError(tg.Name, err.Error())
					break
				}

				msg.Text = "account " + newUser.Username + " created"
			} else {
				currentUser, err := tg.UserService.GetUser(currentAccount.UserId.String())

				if err != nil || currentUser == nil {
					utils.GenerateError(tg.Name, err.Error())
					break
				}

				msg.Text = "account " + currentUser.Username + " already created"
			}
		case "event":
			msg.Text = "start to create event"
			currentAccountId := update.Message.From.ID
			currentNotCompletedEventOfCurrentAccount, ok := tg.NotCompletedEventMap[currentAccountId]

			if ok {
				log.SetPrefix("TG_BOT ")
				log.Println("created event iof current account id:", currentNotCompletedEventOfCurrentAccount)
				msg.Text = "We have not completed event,"
				reflectValue := reflect.ValueOf(currentNotCompletedEventOfCurrentAccount).Elem()
				if isInvalidName := reflectValue.FieldByName("Name").IsZero(); isInvalidName {
					msg.Text += " enter name or info about event"
				} else if isInvalidDate := reflectValue.FieldByName("Date").IsZero(); isInvalidDate {
					msg.Text += " enter date in next format: YYYY-MM-DD"
				}
			} else {
				tg.NotCompletedEventMap[currentAccountId] = &TgEvent{Date: nil}
				msg.Text = "Start to create event, write event name or info"
			}
		case "help":
			msg.Text = "I understand /sayhi and /status."
		default:
			msg.Text = "I don't know that command"
		}

		if len(msg.Text) == 0 {
			msg.Text = "Something went wrong"
		}

		if _, err := tg.Bot.Send(msg); err != nil {
			utils.GenerateError(tg.Name, err.Error())
		}
	}
}
