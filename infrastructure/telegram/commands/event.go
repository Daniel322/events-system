package tg_commands

import (
	"context"
	"events-system/infrastructure/cache"
	"events-system/interfaces"
	"reflect"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func EventCmd(ctx context.Context, msg *tgbotapi.MessageConfig, update tgbotapi.Update) error {
	msg.Text = "start to create event"
	currentAccountId := update.Message.From.ID
	strAccId := strconv.Itoa(int(currentAccountId))

	currentNotCompletedEventOfCurrentAccount, ok := cache.Instance.Get(strAccId)

	if ok {
		msg.Text = "We have not completed event,"
		reflectValue := reflect.ValueOf(currentNotCompletedEventOfCurrentAccount).Elem()
		if isInvalidName := reflectValue.FieldByName("Name").IsZero(); isInvalidName {
			msg.Text += " enter name or info about event"
		} else if isInvalidDate := reflectValue.FieldByName("Date").IsZero(); isInvalidDate {
			msg.Text += " enter date in next format: YYYY-MM-DD"
		}
	} else {
		cache.Instance.Set(strAccId, &interfaces.TgEvent{Date: nil})
		msg.Text = "Start to create event, write event name or info"
	}

	return nil
}
