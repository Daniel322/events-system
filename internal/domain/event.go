package domain

import (
	"encoding/json"
	"events-system/internal/utils"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID           uuid.UUID
	UserId       uuid.UUID
	Info         string
	Date         time.Time
	NotifyLevels []byte
	Providers    []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type EventFactory struct {
	Name string
}

type CreateEventData struct {
	UserId       string
	Info         string
	Date         time.Time
	NotifyLevels []string
	Providers    []string
}

type UpdateEventData struct {
	Info         string
	Date         time.Time
	NotifyLevels []string
	Providers    []string
}

var NOTIFY_LEVELS = []string{"month", "week", "tomorrow", "today"}

const (
	INVALID_INFO          = "invalid info"
	INVALID_NOTIFY_LEVELS = "error on parse notify levels"
	INVALID_PROVIDERS     = "error on parse providers"
)

func NewEventFactory() *EventFactory {
	return &EventFactory{
		Name: "EventFactory",
	}
}

func (ef *EventFactory) Create(data CreateEventData) (*Event, error) {
	var id uuid.UUID = uuid.New()

	parsedUserId, _, err := utils.ParseId(data.UserId)

	if err != nil {
		return nil, utils.GenerateError(ef.Name, err.Error())
	}

	if len(data.Info) == 0 {
		return nil, utils.GenerateError(ef.Name, INVALID_INFO)
	}

	var notifyLevelsForParse []string

	if len(data.NotifyLevels) > 0 {
		notifyLevelsForParse = data.NotifyLevels
	} else {
		notifyLevelsForParse = NOTIFY_LEVELS
	}

	notifyLevels, err := json.Marshal(notifyLevelsForParse)

	if err != nil {
		return nil, utils.GenerateError(ef.Name, INVALID_NOTIFY_LEVELS)
	}

	parsedProviders, err := json.Marshal(data.Providers)

	if err != nil {
		return nil, utils.GenerateError(ef.Name, INVALID_PROVIDERS)
	}

	var event = Event{
		ID:           id,
		UserId:       parsedUserId,
		Info:         data.Info,
		Date:         data.Date,
		NotifyLevels: notifyLevels,
		Providers:    parsedProviders,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return &event, nil
}

func (ef *EventFactory) Update(event *Event, data UpdateEventData) (*Event, error) {
	dataValue := reflect.ValueOf(data).Elem()
	var fields = 0

	if infoField := dataValue.FieldByName("Info"); infoField.IsValid() {
		event.Info = data.Info
		fields++
	}

	if notifyField := dataValue.FieldByName("NotifyLevels"); notifyField.IsValid() {
		parsedNotifyLevels, err := json.Marshal(data.NotifyLevels)

		if err != nil {
			return nil, utils.GenerateError(ef.Name, err.Error())
		}

		event.NotifyLevels = parsedNotifyLevels
		fields++
	}

	if providersField := dataValue.FieldByName("Providers"); providersField.IsValid() {
		parsedProviders, err := json.Marshal(data.Providers)

		if err != nil {
			return nil, utils.GenerateError(ef.Name, err.Error())
		}

		event.Providers = parsedProviders
		fields++
	}

	if dateField := dataValue.FieldByName("Date"); dateField.IsValid() {
		event.Date = data.Date
		fields++
	}

	if fields > 0 {
		event.UpdatedAt = time.Now()
	}

	return event, nil
}
