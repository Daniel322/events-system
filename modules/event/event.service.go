package event_module

import (
	"context"
	"encoding/json"
	"events-system/modules/db"
	"log"
	"time"
)

type CreateEventData struct {
	UserId       string    `json:"user_id"`
	Info         string    `json:"info"`
	Date         time.Time `json:"date"`
	NotifyLevels []string  `json:"notify_levels"`
	Providers    []string  `json:"providers"`
}

type UpdateEventData struct {
	Info string    `json:"info"`
	Date time.Time `json:"date"`
}

func GetUserEvents(userId string) (*[]Event, error) {
	var events []Event

	result := db.Connection.Table("events").Where("user_id = ?", userId).Take(&events)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &events, nil

}

func UpdateEvent(id string, data UpdateEventData) (*Event, error) {
	var event Event
	result := db.Connection.Model(&event).Where("id = ?", id).Updates(data)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &event, nil
}

func CreateEvent(data CreateEventData, currentContext context.Context) (*Event, error) {
	jsonProviders, err := json.Marshal(data.Providers)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	event := Event{
		UserId:    data.UserId,
		Info:      data.Info,
		Date:      data.Date,
		Providers: jsonProviders,
	}

	result := db.Connection.Create(&event)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &event, nil
}

func DeleteEvent(id string, operationContext context.Context) (bool, error) {
	result := db.Connection.Table("events").Delete(&Event{}, id)

	if result.Error != nil {
		log.Fatal(result.Error)
		return false, result.Error
	}

	return true, nil
}
