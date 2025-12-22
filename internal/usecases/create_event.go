package usecases

import (
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"time"
)

func (usecase *InternalUseCases) generateTimesForTasks(
	eventDate time.Time,
	providers entities.JsonField,
) []dto.TaskSliceEvent {
	today := time.Now()
	todayYear := today.Year()
	tasks := make([]dto.TaskSliceEvent, 0)

	for key, value := range TASKS_TYPES {
		currentEventInThatYear := time.Date(
			todayYear,
			eventDate.Month(),
			eventDate.Day(),
			eventDate.Hour(),
			eventDate.Minute(),
			eventDate.Second(),
			eventDate.Nanosecond(),
			eventDate.Location(),
		).Add(-value)

		if currentEventInThatYear.Compare(today) == -1 {
			currentEventInThatYear = currentEventInThatYear.AddDate(1, 0, 0)
		}

		for _, provider := range providers {
			tasks = append(tasks, dto.TaskSliceEvent{
				Type:     key,
				Date:     currentEventInThatYear,
				Provider: provider,
			})
		}
	}

	return tasks
}

func (usecase *InternalUseCases) CreateEvent(data dto.CreateEventDTO) (*dto.OutputEvent, error) {
	transaction := usecase.BaseRepository.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	event, err := usecase.EventService.Create(
		dto.CreateEventData{
			UserId:    data.UserId,
			Info:      data.Info,
			Date:      data.Date,
			Providers: data.Providers,
		},
		transaction,
	)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError("CreateEvent", err.Error())
	}

	tasks_dates := usecase.generateTimesForTasks(event.Date, data.Providers)

	tasks := make([]entities.Task, 0)

	for _, value := range tasks_dates {
		task, err := usecase.TaskService.Create(
			dto.CreateTaskData{
				EventId:   event.ID,
				AccountId: data.AccountId,
				Type:      value.Type,
				Date:      value.Date,
				Provider:  value.Provider,
			},
			transaction,
		)

		if err != nil {
			transaction.Rollback()
			return nil, utils.GenerateError("CreateEvent", err.Error())
		}

		tasks = append(tasks, *task)
	}

	if trRes := transaction.Commit(); trRes.Error != nil {
		return nil, utils.GenerateError("CreateEvent", trRes.Error.Error())
	}

	return &dto.OutputEvent{
		ID:           event.ID,
		UserId:       event.UserId,
		Info:         event.Info,
		Date:         event.Date,
		NotifyLevels: event.NotifyLevels,
		Providers:    event.Providers,
		CreatedAt:    event.CreatedAt,
		UpdatedAt:    event.UpdatedAt,
		Tasks:        tasks,
	}, nil
}
