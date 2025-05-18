package event_module

import (
	"context"
	"encoding/json"
	"events-system/modules/db"
	"log"
)

type CreateEventData struct {
	UserId       string   `json:"user_id"`
	Info         string   `json:"info"`
	Date         string   `json:"date"`
	NotifyLevels []string `json:"notify_levels"`
	Providers    []string `json:"providers"`
}

type UpdateEventData struct {
}

func GetUserEvents(userId string) (*[]Event, error) {
	query := "SELECT * FROM events WHERE user_id = $1"
	result, err := db.BaseQuery[[]Event](context.Background(), query, userId)
	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func CreateEvent(data CreateEventData, currentContext context.Context) (*Event, error) {
	query := "INSERT INTO events (user_id, info, date, providers) VALUES ($1, $2, $3, $4) RETURNING *"

	// jsonNotify, err := json.Marshal(data.NotifyLevels)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	jsonProviders, err := json.Marshal(data.Providers)
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.BaseQuery[Event](
		currentContext,
		query,
		data.UserId,
		data.Info,
		data.Date,
		jsonProviders,
	)
	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func DeleteEvent(id string, operationContext context.Context) (bool, error) {
	query := "DELETE FROM events WHERE id = $1"
	_, err := db.Connection.Exec(operationContext, query, id)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, err
}
