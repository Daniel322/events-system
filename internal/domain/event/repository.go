package event

import (
	"context"
	"events-system/interfaces"
	"events-system/internal/components"
)

type EventsRepo struct {
	components.Factory
}

func NewEventsRepo(repo interfaces.RepositoryV2) *EventsRepo {
	return &EventsRepo{
		Factory: *components.NewFactory("Event", repo),
	}
}

func (r EventsRepo) Save(ctx context.Context, value interface{}) error {
	ctx = context.WithValue(ctx, "tableName", "events")

	return r.Repository.Save(ctx, value)
}
