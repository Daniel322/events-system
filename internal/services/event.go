package services

// import (
// 	"events-system/interfaces"
// 	"events-system/internal/dto"
// 	entities "events-system/internal/entity"
// 	dependency_container "events-system/pkg/di"
// 	"events-system/pkg/repository"
// 	"events-system/pkg/utils"
// )

// type EventService struct {
// 	Name string
// }

// func NewEventService() *EventService {
// 	service := &EventService{
// 		Name: "EventService",
// 	}

// 	dependency_container.Container.Add("eventService", service)

// 	return service
// }

// func (es *EventService) CreateEvent(data dto.CreateEventDTO) (*dto.OutputEvent, error) {
// 	dependencies := []string{"eventFactory", "taskService"}
// 	diServices, err := dependency_container.Container.MultiGet(dependencies)

// 	if err != nil {
// 		return nil, utils.GenerateError(es.Name, err.Error())
// 	}

// 	eventFactory := diServices["eventFactory"]

// 	taskService := diServices["taskService"]

// 	transaction := repository.CreateTransaction()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			transaction.Rollback()
// 		}
// 	}()

// 	event, err := eventFactory.(interfaces.EventFactory).Create(entities.CreateEventData{
// 		UserId:       data.UserId,
// 		Info:         data.Info,
// 		Date:         data.Date,
// 		Providers:    data.Providers,
// 		NotifyLevels: []string{"month", "week", "tomorrow", "today"},
// 	})

// 	if err != nil {
// 		return nil, utils.GenerateError(es.Name, err.Error())
// 	}

// 	event, err = repository.Create(repository.Events, *event, transaction)

// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, utils.GenerateError(es.Name, err.Error())
// 	}

// 	timesForTask := taskService.(interfaces.TaskService).GenerateTimesForTasks(data.Date)

// 	tasks := make([]entities.Task, 0)

// 	for _, timeValue := range timesForTask {
// 		uuidV, _, err := utils.ParseId(data.AccountId)

// 		if err != nil {
// 			transaction.Rollback()
// 			return nil, utils.GenerateError(es.Name, err.Error())
// 		}

// 		task, err := taskService.(interfaces.TaskService).Create(entities.CreateTaskData{
// 			EventId:   event.ID,
// 			AccountId: uuidV,
// 			Date:      timeValue.Date,
// 			Type:      timeValue.Type,
// 			Provider:  "telegram",
// 		}, transaction)

// 		if err != nil {
// 			transaction.Rollback()
// 			return nil, utils.GenerateError(es.Name, err.Error())
// 		}

// 		tasks = append(tasks, *task)
// 	}

// 	transaction.Commit()

// 	return &dto.OutputEvent{
// 		ID:           event.ID,
// 		UserId:       event.UserId,
// 		Info:         event.Info,
// 		Date:         event.Date,
// 		NotifyLevels: event.NotifyLevels,
// 		Providers:    event.Providers,
// 		CreatedAt:    event.CreatedAt,
// 		UpdatedAt:    event.UpdatedAt,
// 		Tasks:        tasks,
// 	}, nil
// }
