package tg_handlers

import (
	"context"
	"events-system/infrastructure/cache"
	"events-system/interfaces"
	"events-system/internal/application/commands"
	"events-system/internal/application/queries"
	"events-system/internal/domain/account"
	"events-system/pkg/utils"
	"reflect"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageHandler(ctx context.Context, msg *tgbotapi.MessageConfig, update tgbotapi.Update) error {
	currentAccountId := update.Message.From.ID
	strAccId := strconv.Itoa(int(currentAccountId))
	t, err := account.NewAccountType("telegram")

	if err != nil {
		return utils.GenerateError("MessageHandler", err.Error())
	}

	checkAccState, err := queries.NewCheckAccountState(strAccId, t)

	if err != nil {
		return utils.GenerateError("MessageHandler", err.Error())
	}

	currentAcc, err := queries.CheckAccount.Run(ctx, *checkAccState)

	if err != nil {
		return utils.GenerateError("MessageHandler", err.Error())
	}

	if currentAcc == nil {
		msg.Text = "need to create account, use /start command"
		return nil
	}

	currentEvent, ok := cache.Instance.Get(strAccId)

	if !ok {
		msg.Text = "you dont have uncompleted events, use /event command for start to create event or /help for get list of available commands"
		return nil
	} else {
		reflectValue := reflect.ValueOf(currentEvent).Elem()
		if isZeroName := reflectValue.FieldByName("Name").IsZero(); isZeroName {
			cache.Instance.Set(strAccId, &interfaces.TgEvent{Date: nil, Name: update.Message.Text})
			msg.Text = "added name, now add Date in next format YYYY-MM-DD"
			return nil
		} else {
			timeVar, err := time.Parse("2006-01-02", update.Message.Text)

			if err != nil {
				err = utils.GenerateError("MessageHandler", err.Error())
				msg.Text = err.Error()
				return nil
			}

			eventState, err := commands.CreateEvent.Validate(commands.CreateEventData{
				Info:         currentEvent.(*interfaces.TgEvent).Name,
				Date:         timeVar,
				AccId:        currentAcc.ID,
				UserId:       currentAcc.UserId,
				Providers:    []string{"telegram"},
				NotifyLevels: []string{"today", "tomorrow", "week", "month"},
			})

			if err != nil {
				err = utils.GenerateError("MessageHandler", err.Error())
				msg.Text = err.Error()
				return nil
			}

			event, err := commands.CreateEvent.Run(ctx, eventState)

			if err != nil {
				err = utils.GenerateError("MessageHandler", err.Error())
				msg.Text = err.Error()
				return nil
			}

			msg.Text = "event " + event.ToPlain().Info + " with next date:" + event.ToPlain().Date.Format("2006-01-02") + " created!"
			cache.Instance.Remove(strAccId)
			return nil
		}
	}
}
