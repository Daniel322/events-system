package domain

import (
	"encoding/json"
	"errors"
	"events-system/internal/utils"
	"fmt"
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

func NewEventFactory() *EventFactory {
	return &EventFactory{
		Name: "EventFactory",
	}
}

func (ef *EventFactory) Create(data CreateEventData) (*Event, error) {
	var id uuid.UUID = uuid.New()

	parsedUserId, _, err := utils.ParseId(data.UserId)

	if err != nil {
		fmt.Println("Invalid userId type")
		return nil, err
	}

	if len(data.Info) == 0 {
		return nil, errors.New("invalid info")
	}

	var notifyLevelsForParse []string

	if len(data.NotifyLevels) > 0 {
		notifyLevelsForParse = data.NotifyLevels
	} else {
		notifyLevelsForParse = NOTIFY_LEVELS
	}

	notifyLevels, err := json.Marshal(notifyLevelsForParse)

	if err != nil {
		fmt.Println("Error on parse notify levels")
		return nil, err
	}

	parsedProviders, err := json.Marshal(data.Providers)

	if err != nil {
		fmt.Println("Error on parse providers")
		return nil, err
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
	event.UpdatedAt = time.Now()

	if infoField := dataValue.FieldByName("Info"); infoField.IsValid() {
		event.Info = data.Info
	}

	if notifyField := dataValue.FieldByName("NotifyLevels"); notifyField.IsValid() {
		parsedNotifyLevels, err := json.Marshal(data.NotifyLevels)

		if err != nil {
			fmt.Println("Error on parse notify levels")
			return nil, err
		}

		event.NotifyLevels = parsedNotifyLevels
	}

	if providersField := dataValue.FieldByName("Providers"); providersField.IsValid() {
		parsedProviders, err := json.Marshal(data.Providers)

		if err != nil {
			fmt.Println("Error on parse providers")
			return nil, err
		}

		event.Providers = parsedProviders
	}

	if dateField := dataValue.FieldByName("Date"); dateField.IsValid() {
		event.Date = data.Date
	}

	return event, nil
}
