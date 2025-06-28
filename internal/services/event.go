package services

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/repositories"
)

type EventService struct {
	Name            string
	DB              *db.Database
	eventRepository *repositories.Repository[domain.Event, domain.CreateEventData, domain.UpdateEventData]
	taskRepository  *repositories.Repository[domain.Task, domain.CreateTaskData, domain.UpdateTaskData]
}

func NewEventService(
	db *db.Database,
	eventRepository *repositories.Repository[domain.Event, domain.CreateEventData, domain.UpdateEventData],
	taskRepository *repositories.Repository[domain.Task, domain.CreateTaskData, domain.UpdateTaskData],
) *EventService {
	return &EventService{
		Name:            "EventService",
		DB:              db,
		eventRepository: eventRepository,
		taskRepository:  taskRepository,
	}
}
