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
	EventService         services.IEventService
	NotCompletedEventMap map[int64]*TgEvent
}

func NewTgBotProvider(
	token string,
	userService services.IUserService,
	accService services.IAccountService,
	eventService services.IEventService,
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

		if !update.Message.IsCommand() {
			log.Println(tg.NotCompletedEventMap)
			currentEvent, isCurrentEventExist := tg.NotCompletedEventMap[update.Message.From.ID]
			log.Println(currentEvent, isCurrentEventExist)
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
					log.Println("timeVar", timeVar.String(), timeVar)
					// strAccId := strconv.Itoa(int(update.Message.From.ID))
					currentEvent.Date = &timeVar

					event, err := tg.EventService.CreateEvent(services.CreateEventData{
						AccountId: currentAcc.ID.String(),
						Date:      *currentEvent.Date,
						Info:      currentEvent.Name,
						UserId:    currentAcc.UserId.String(),
					})

					if err != nil {
						msg.Text = err.Error()
						tg.Bot.Send(msg)
						continue
					}

					log.Println("EVENT:", event)
					log.Println("err:", err)

					msg.Text = "event " + currentEvent.Name + " with next date:" + currentEvent.Date.Format("2006-01-02") + " created!"
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
		}

		if len(msg.Text) == 0 {
			msg.Text = "Something went wrong"
		}

		if _, err := tg.Bot.Send(msg); err != nil {
			utils.GenerateError(tg.Name, err.Error())
		}
	}
}
