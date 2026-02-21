package event

import (
	"context"
	"events-system/interfaces"
	"events-system/internal/components"
	"events-system/pkg/utils"
)

type EventsRepo struct {
	components.Factory
}

func NewEventsRepo(repo interfaces.Repository) *EventsRepo {
	return &EventsRepo{
		Factory: *components.NewFactory("Event", repo),
	}
}

func (r EventsRepo) Save(ctx context.Context, value interface{}) error {
	ctx = context.WithValue(ctx, "tableName", "events")

	return r.Repository.Save(ctx, value)
}

func (r EventsRepo) FindOne(ctx context.Context, options map[string]interface{}) (*Plain, error) {
	events := new([]Plain)
	ctx = context.WithValue(ctx, "tableName", "tasks")
	ctx = context.WithValue(ctx, "ptr", events)
	err := r.Repository.Find(ctx, options)

	if err != nil {
		return nil, utils.GenerateError("TaskRepo FindOne", err.Error())
	}
	if len(*events) == 0 {
		return nil, nil
	}

	event := (*events)[0]

	return &event, nil
}
