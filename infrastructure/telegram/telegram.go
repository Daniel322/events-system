package telegram

import (
	"context"
	"events-system/infrastructure/config"
	tg_commands "events-system/infrastructure/telegram/commands"
	"events-system/pkg/utils"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgEvent struct {
	Name string
	Date *time.Time
}

type TgBotProvider struct {
	Logger               *log.Logger
	Bot                  *tgbotapi.BotAPI
	NotCompletedEventMap map[int64]*TgEvent
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
		Logger:               logger,
		Bot:                  bot,
		NotCompletedEventMap: make(map[int64]*TgEvent, 10),
	}

	return nil
}

// func (tg *TgBotProvider) NewMessage(id int64) tgbotapi.MessageConfig {
// 	return tgbotapi.NewMessage(id, "")
// }

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
			// here will be long process handlers
		} else {
			switch update.Message.Command() {
			// here will be our cmd handlers
			case "help":
				tg_commands.HelpCmd(ctx, &msg)
			case "start":
				tg_commands.StartCmd(&msg, update)
			}
		}

		// if !update.Message.IsCommand() {
		// 	currentAcc, err := tg.Service.CheckTGAccount(update.Message.Chat.ID)

		// 	if err != nil {
		// 		msg.Text = err.Error()
		// 		tg.Bot.Send(msg)
		// 		continue
		// 	}

		// 	if currentAcc == nil {
		// 		msg.Text = "need to create account, use /start command"
		// 		tg.Bot.Send(msg)
		// 		continue
		// 	}

		// 	currentEvent, isCurrentEventExist := tg.NotCompletedEventMap[update.Message.From.ID]

		// 	if !isCurrentEventExist {
		// 		msg.Text = "you dont have uncompleted events, use /event command for start to create event or /help for get list of available commands"
		// 		tg.Bot.Send(msg)
		// 		continue
		// 	} else {
		// 		reflectValue := reflect.ValueOf(currentEvent).Elem()
		// 		if isZeroName := reflectValue.FieldByName("Name").IsZero(); isZeroName {
		// 			currentEvent.Name = update.Message.Text
		// 			msg.Text = "added name, now add Date in next format YYYY-MM-DD"
		// 			tg.Bot.Send(msg)
		// 			continue
		// 		} else {
		// 			timeVar, err := time.Parse("2006-01-02", update.Message.Text)
		// 			if err != nil {
		// 				err = utils.GenerateError(tg.Name, err.Error())
		// 				msg.Text = err.Error()
		// 				tg.Bot.Send(msg)
		// 				continue
		// 			}

		// 			// strAccId := strconv.Itoa(int(update.Message.From.ID))
		// 			currentEvent.Date = &timeVar
		// 			event, err := tg.Service.CreateEvent(dto.CreateEventDTO{
		// 				AccountId: currentAcc.ID,
		// 				Date:      *currentEvent.Date,
		// 				Info:      currentEvent.Name,
		// 				UserId:    currentAcc.UserId,
		// 				Providers: entities.JsonField{"telegram"},
		// 			})

		// 			if err != nil {
		// 				msg.Text = err.Error()
		// 				tg.Bot.Send(msg)
		// 				continue
		// 			}

		// 			msg.Text = "event " + event.Info + " with next date:" + event.Date.Format("2006-01-02") + " created!"
		// 			delete(tg.NotCompletedEventMap, update.Message.From.ID)
		// 			tg.Bot.Send(msg)
		// 			continue
		// 		}
		// 	}
		// } else {
		// 	switch update.Message.Command() {
		// 	case "start":
		// 		accountId := update.Message.From.ID
		// 		currentAccount, err := tg.Service.CheckTGAccount(accountId)

		// 		if err != nil {
		// 			utils.GenerateError(tg.Name, err.Error())
		// 			break
		// 		}

		// 		if currentAccount == nil {
		// 			strAccId := strconv.Itoa(int(update.Message.From.ID))

		// 			newUser, err := tg.Service.CreateUser(dto.CreateUserInput{
		// 				Username:  update.Message.From.UserName,
		// 				AccountId: strAccId,
		// 				Type:      entities.AccountType(1),
		// 			})

		// 			if err != nil {
		// 				utils.GenerateError(tg.Name, err.Error())
		// 				break
		// 			}

		// 			msg.Text = "account " + newUser.Username + " created"
		// 		} else {
		// 			currentUser, err := tg.Service.GetUser(currentAccount.UserId.String())

		// 			if err != nil || currentUser == nil {
		// 				utils.GenerateError(tg.Name, err.Error())
		// 				break
		// 			}

		// 			msg.Text = "account " + currentUser.Username + " already created"
		// 		}
		// 	case "event":
		// 		msg.Text = "start to create event"
		// 		currentAccountId := update.Message.From.ID
		// 		currentNotCompletedEventOfCurrentAccount, ok := tg.NotCompletedEventMap[currentAccountId]

		// 		if ok {
		// 			log.SetPrefix("TG_BOT ")
		// 			log.Println("created event for current account id:", currentNotCompletedEventOfCurrentAccount)
		// 			msg.Text = "We have not completed event,"
		// 			reflectValue := reflect.ValueOf(currentNotCompletedEventOfCurrentAccount).Elem()
		// 			if isInvalidName := reflectValue.FieldByName("Name").IsZero(); isInvalidName {
		// 				msg.Text += " enter name or info about event"
		// 			} else if isInvalidDate := reflectValue.FieldByName("Date").IsZero(); isInvalidDate {
		// 				msg.Text += " enter date in next format: YYYY-MM-DD"
		// 			}
		// 		} else {
		// 			tg.NotCompletedEventMap[currentAccountId] = &TgEvent{Date: nil}
		// 			msg.Text = "Start to create event, write event name or info"
		// 		}
		// 	case "help":
		// 		msg.Text = "I understand /start and /event."
		// 	default:
		// 		msg.Text = "I don't know that command"
		// 	}
		// }

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
