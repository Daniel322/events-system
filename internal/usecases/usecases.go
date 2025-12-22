package usecases

import (
	"events-system/interfaces"
	"time"
)

type InternalUseCases struct {
	BaseRepository interfaces.BaseRepository
	UserService    interfaces.UserService
	AccountService interfaces.AccountService
	EventService   interfaces.EventService
	TaskService    interfaces.TaskService
}

var TASKS_TYPES = map[string]time.Duration{
	"today":    0,
	"tomorrow": time.Hour * 24,
	"week":     time.Hour * 168,
	"month":    time.Hour * 720,
}

func NewInternalUseCases(
	repository interfaces.BaseRepository,
	user_service interfaces.UserService,
	acc_service interfaces.AccountService,
	event_service interfaces.EventService,
	task_service interfaces.TaskService,
) *InternalUseCases {
	return &InternalUseCases{
		BaseRepository: repository,
		UserService:    user_service,
		AccountService: acc_service,
		EventService:   event_service,
		TaskService:    task_service,
	}
}
