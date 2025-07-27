package telegram

import (
	"events-system/internal/interfaces"
	"events-system/internal/structs"
	"events-system/internal/utils"
	"log"
	"reflect"
	"strconv"
	"strings"
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
	UserService          interfaces.IUserService
	AccountService       interfaces.IAccountService
	EventService         interfaces.IEventService
	NotCompletedEventMap map[int64]*TgEvent
}

func NewTgBotProvider(
	token string,
	userService interfaces.IUserService,
	accService interfaces.IAccountService,
	eventService interfaces.IEventService,
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
		EventService:         eventService,
		NotCompletedEventMap: make(map[int64]*TgEvent, 10),
	}, nil
}

func (tg *TgBotProvider) NewMessage(id int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(id, "")
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

		if !update.Message.IsCommand() {
			currentAcc, err := tg.AccountService.CheckAccount(update.Message.Chat.ID)

			if err != nil {
				msg.Text = err.Error()
				tg.Bot.Send(msg)
				continue
			}

			if currentAcc == nil {
				msg.Text = "need to create account, use /start command"
				tg.Bot.Send(msg)
				continue
			}

			currentEvent, isCurrentEventExist := tg.NotCompletedEventMap[update.Message.From.ID]

			if !isCurrentEventExist {
				msg.Text = "you dont have uncompleted events, use /event command for start to create event or /help for get list of available commands"
				tg.Bot.Send(msg)
				continue
			} else {
				reflectValue := reflect.ValueOf(currentEvent).Elem()
				if isZeroName := reflectValue.FieldByName("Name").IsZero(); isZeroName {
					currentEvent.Name = update.Message.Text
					msg.Text = "added name, now add Date in next format YYYY-MM-DD"
					tg.Bot.Send(msg)
					continue
				} else {
					timeVar, err := time.Parse("2006-01-02", update.Message.Text)
					if err != nil {
						err = utils.GenerateError(tg.Name, err.Error())
						msg.Text = err.Error()
						tg.Bot.Send(msg)
						continue
					}

					// strAccId := strconv.Itoa(int(update.Message.From.ID))
					currentEvent.Date = &timeVar

					event, err := tg.EventService.CreateEvent(structs.CreateEventData{
						AccountId: currentAcc.ID.String(),
						Date:      *currentEvent.Date,
						Info:      currentEvent.Name,
						UserId:    currentAcc.UserId.String(),
						Providers: []byte(strings.Join([]string{"telegram"}, " ")),
					})

					if err != nil {
						msg.Text = err.Error()
						tg.Bot.Send(msg)
						continue
					}

					msg.Text = "event " + event.Info + " with next date:" + event.Date.Format("2006-01-02") + " created!"
					delete(tg.NotCompletedEventMap, update.Message.From.ID)
					tg.Bot.Send(msg)
					continue
				}
			}
		} else {
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

					newUser, err := tg.UserService.CreateUser(structs.CreateUserData{
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
					log.Println("created event for current account id:", currentNotCompletedEventOfCurrentAccount)
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
		}

		if len(msg.Text) == 0 {
			msg.Text = "Something went wrong"
		}

		if _, err := tg.Bot.Send(msg); err != nil {
			utils.GenerateError(tg.Name, err.Error())
		}
	}
}

func (tg *TgBotProvider) SendMsg(chatId int64, text string) {
	msg := tgbotapi.NewMessage(chatId, text)

	tg.Bot.Send(msg)
}
