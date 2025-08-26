package services

import (
	"events-system/interfaces"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	dependency_container "events-system/pkg/di"
	"events-system/pkg/repository"
	"events-system/pkg/utils"
	"log"
	"time"
)

type EventService struct {
	Name string
}

type TaskSliceEvent struct {
	Date time.Time
	Type string
}

func NewEventService() *EventService {
	service := &EventService{
		Name: "EventService",
	}

	dependency_container.Container.Add("eventService", service)

	return service
}

func (es *EventService) CreateEvent(data dto.CreateEventDTO) (*dto.OutputEvent, error) {
	eventFactory, err := dependency_container.Container.Get("eventFactory")

	if err != nil {
		return nil, utils.GenerateError(es.Name, err.Error())
	}

	taskService, err := dependency_container.Container.Get("taskService")

	if err != nil {
		return nil, utils.GenerateError(es.Name, err.Error())
	}

	transaction := repository.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	event, err := eventFactory.(interfaces.EventFactory).Create(entities.CreateEventData{
		UserId:       data.UserId,
		Info:         data.Info,
		Date:         data.Date,
		Providers:    data.Providers,
		NotifyLevels: []string{"month", "week", "tomorrow", "today"},
	})

	if err != nil {
		return nil, utils.GenerateError(es.Name, err.Error())
	}

	event, err = repository.Create(repository.Events, *event, transaction)

	log.SetPrefix("event service info")
	log.Println(event)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(es.Name, err.Error())
	}

	timesForTask := make([]TaskSliceEvent, 0)

	tasks := make([]entities.Task, 0)

	today := time.Now()
	todayYear := today.Year()
	eventDateYear := data.Date.Year()
	currentEventInThatYear := data.Date
	// TODO: check flow and fix bug with next case: если создать евент с таском в текущий день, таск создастся на следующий год
	if eventDateYear < todayYear {
		currentEventInThatYear = time.Date(
			todayYear,
			data.Date.Month(),
			data.Date.Day(),
			data.Date.Hour(),
			data.Date.Minute(),
			data.Date.Second(),
			data.Date.Nanosecond(),
			data.Date.Location(),
		)
		// if event in that year before today
		if currentEventInThatYear.Compare(today) == -1 {
			currentEventInThatYear = time.Date(
				todayYear+1,
				data.Date.Month(),
				data.Date.Day(),
				data.Date.Hour(),
				data.Date.Minute(),
				data.Date.Second(),
				data.Date.Nanosecond(),
				data.Date.Location(),
			)
		}
	}
	timesForTask = append(timesForTask, TaskSliceEvent{Date: currentEventInThatYear, Type: "today"})
	timesForTask = append(timesForTask, TaskSliceEvent{Date: currentEventInThatYear.Add(-(time.Hour * 24)), Type: "tomorrow"})
	timesForTask = append(timesForTask, TaskSliceEvent{Date: currentEventInThatYear.Add(-(time.Hour * 24 * 7)), Type: "week"})
	timesForTask = append(timesForTask, TaskSliceEvent{Date: currentEventInThatYear.Add(-(time.Hour * 24 * 30)), Type: "month"})

	for _, timeValue := range timesForTask {
		log.Println("task create start")
		uuidV, _, err := utils.ParseId(data.AccountId)

		if err != nil {
			log.Println(err)
			transaction.Rollback()
			return nil, utils.GenerateError(es.Name, err.Error())
		}

		task, err := taskService.(interfaces.TaskService).Create(entities.CreateTaskData{
			EventId:   event.ID,
			AccountId: uuidV,
			Date:      timeValue.Date,
			Type:      timeValue.Type,
			Provider:  "telegram",
		}, transaction)

		if err != nil {
			transaction.Rollback()
			return nil, utils.GenerateError(es.Name, err.Error())
		}

		tasks = append(tasks, *task)
	}

	transaction.Commit()

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
