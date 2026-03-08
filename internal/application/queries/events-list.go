package queries

import (
	"context"
	"events-system/internal/domain/event"
	"events-system/pkg/utils"
	"log"
)

type IEventsList struct {
	logger     *log.Logger
	repository *event.EventsRepo
}

var EventsList *IEventsList

func InitEventsList() {
	var logger = log.New(log.Writer(), "EventsList ", log.LstdFlags)

	EventsList = &IEventsList{
		repository: event.Repository,
		logger:     logger,
	}
}

func (this IEventsList) Run(ctx context.Context, userId string) (*[]event.Plain, error) {
	options := make(map[string]interface{})
	options["user_id"] = userId
	result := new([]event.Plain)
	ctx = context.WithValue(ctx, "tableName", "events")
	ctx = context.WithValue(ctx, "ptr", result)

	err := this.repository.Repository.Find(ctx, options)

	if err != nil {
		return nil, utils.GenerateError("EventsList.Run", err.Error())
	}

	return result, nil
}
