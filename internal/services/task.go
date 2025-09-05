package services

import (
	"events-system/infrastructure/providers/db"
	"events-system/interfaces"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"slices"
	"time"

	"github.com/google/uuid"
)

type TaskService struct {
	Name       string
	Repository interfaces.Repository[entities.Task]
}

const (
	DATE_IS_REQUIRED      = "date is required"
	INVALID_TASK_TYPE     = "type is invalid"
	INVALID_TASK_PROVIDER = "provider is invalid"
)

func NewTaskService(repository interfaces.Repository[entities.Task]) *TaskService {
	return &TaskService{
		Name:       "TaskService",
		Repository: repository,
	}
}

func (service *TaskService) checkDate(value time.Time) error {
	if value.IsZero() {
		return utils.GenerateError(service.Name, INVALID_DATE)
	}
	return nil
}

func (service *TaskService) checkContainsOfSupportedvalues(value string, slice []string, err string) error {
	if isContains := slices.Contains(slice, value); !isContains {
		return utils.GenerateError(service.Name, err)
	}

	return nil
}

func (service *TaskService) Find(options map[string]interface{}) (*[]entities.Task, error) {
	results, err := service.Repository.Find(options)

	return results, err
}

func (service *TaskService) Delete(id string, transaction db.DatabaseInstance) (bool, error) {
	result, err := service.Repository.Destroy(id, transaction)

	return result, err
}

func (service *TaskService) Create(
	data dto.CreateTaskData,
	transaction db.DatabaseInstance,
) (*entities.Task, error) {
	var id uuid.UUID = uuid.New()

	if err := uuid.Validate(data.AccountId.String()); err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if err := uuid.Validate(data.EventId.String()); err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	err := service.checkDate(data.Date)

	if err != nil {
		return nil, err
	}

	err = service.checkContainsOfSupportedvalues(data.Type, entities.SUPPORTED_TYPES, INVALID_TASK_TYPE)

	if err != nil {
		return nil, err
	}

	err = service.checkContainsOfSupportedvalues(data.Provider, entities.SUPPORTED_PROVIDERS, INVALID_TASK_PROVIDER)

	if err != nil {
		return nil, err
	}

	task := &entities.Task{
		ID:        id,
		EventId:   data.EventId,
		AccountId: data.AccountId,
		Type:      data.Type,
		Provider:  data.Provider,
		Date:      data.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	task, err = service.Repository.Save(*task, transaction)

	return task, err
}

func (service *TaskService) Update(
	id string,
	date time.Time,
	transaction db.DatabaseInstance,
) (*entities.Task, error) {
	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	tasks, err := service.Repository.Find(findOptions)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if len(*tasks) == 0 {
		return nil, utils.GenerateError(service.Name, "current task with id "+id+" not found")
	}

	currentTask := (*tasks)[0]

	if IsInvalidDate := service.checkDate(date); IsInvalidDate == nil {
		currentTask.Date = date
		currentTask.UpdatedAt = time.Now()

		updatedTask, err := service.Repository.Save(currentTask, transaction)

		return updatedTask, err
	}

	return &currentTask, nil
}

// ------ move to use cases --------

// func (service *TaskService) GenerateTimesForTasks(eventDate time.Time) []entities.TaskSliceEvent {
// 	today := time.Now()
// 	todayYear := today.Year()
// 	eventDateYear := eventDate.Year()
// 	currentEventInThatYear := eventDate
// 	tasks := make([]entities.TaskSliceEvent, 0)
// 	// TODO: check flow and fix bug with next case: если создать евент с таском в текущий день, таск создастся на следующий год
// 	if eventDateYear < todayYear {
// 		currentEventInThatYear = time.Date(
// 			todayYear,
// 			eventDate.Month(),
// 			eventDate.Day(),
// 			eventDate.Hour(),
// 			eventDate.Minute(),
// 			eventDate.Second(),
// 			eventDate.Nanosecond(),
// 			eventDate.Location(),
// 		)
// 		// if event in that year before today
// 		if currentEventInThatYear.Compare(today) == -1 {
// 			currentEventInThatYear = time.Date(
// 				todayYear+1,
// 				eventDate.Month(),
// 				eventDate.Day(),
// 				eventDate.Hour(),
// 				eventDate.Minute(),
// 				eventDate.Second(),
// 				eventDate.Nanosecond(),
// 				eventDate.Location(),
// 			)
// 		}
// 	}
// 	tasks = append(tasks, entities.TaskSliceEvent{Date: currentEventInThatYear, Type: "today"})
// 	tasks = append(tasks, entities.TaskSliceEvent{Date: currentEventInThatYear.Add(-(time.Hour * 24)), Type: "tomorrow"})
// 	tasks = append(tasks, entities.TaskSliceEvent{Date: currentEventInThatYear.Add(-(time.Hour * 24 * 7)), Type: "week"})
// 	tasks = append(tasks, entities.TaskSliceEvent{Date: currentEventInThatYear.Add(-(time.Hour * 24 * 30)), Type: "month"})

// 	return tasks
// }

// func (service *TaskService) Create(
// 	data entities.CreateTaskData,
// 	transaction db.DatabaseInstance,
// ) (*entities.Task, error) {
// 	taskFactory, err := dependency_container.Container.Get("taskFactory")

// 	if err != nil {
// 		return nil, utils.GenerateError(service.Name, err.Error())
// 	}

// 	task, err := taskFactory.(interfaces.TaskFactory).Create(data)

// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, utils.GenerateError(service.Name, err.Error())
// 	}

// 	task, err = repository.Create(repository.Tasks,
// 		*task,
// 		transaction,
// 	)

// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, utils.GenerateError(service.Name, err.Error())
// 	}

// 	return task, nil
// }

// func (ts *TaskService) GetListOfTodayTasks() (*[]entities.Task, error) {
// 	var options = make(map[string]interface{})
// 	options["date"] = time.Now().Format("2006-01-02")
// 	tasks, err := repository.GetList[entities.Task](repository.Tasks, options)

// 	if err != nil {
// 		return nil, utils.GenerateError(ts.Name, err.Error())
// 	}

// 	return tasks, nil
// }

// func (ts *TaskService) ExecTaskAndGenerateNew(taskId string) (*InfoAboutTaskForTgProvider, error) {
// 	currentTask, err := repository.GetById[entities.Task](repository.Tasks, taskId)

// 	if err != nil {
// 		return nil, utils.GenerateError(ts.Name, err.Error())
// 	}

// 	_, strEventId, err := utils.ParseId(currentTask.EventId)

// 	if err != nil {
// 		return nil, utils.GenerateError(ts.Name, err.Error())
// 	}

// 	_, strAccId, err := utils.ParseId(currentTask.AccountId)

// 	if err != nil {
// 		return nil, utils.GenerateError(ts.Name, err.Error())
// 	}

// 	currentEvent, err := repository.GetById[entities.Event](repository.Events, strEventId)

// 	if err != nil {
// 		return nil, utils.GenerateError(ts.Name, err.Error())
// 	}

// 	currentAcc, err := repository.GetById[entities.Account](repository.Accounts, strAccId)

// 	if err != nil {
// 		return nil, utils.GenerateError(ts.Name, err.Error())
// 	}

// 	chatId, err := strconv.ParseInt(currentAcc.AccountId, 10, 64)

// 	if err != nil {
// 		return nil, utils.GenerateError(ts.Name, err.Error())
// 	}

// 	transaction := repository.CreateTransaction()

// 	ok, err := repository.Delete[entities.Task](repository.Tasks, currentTask.ID.String(), transaction)

// 	defer func() {
// 		if r := recover(); r != nil {
// 			transaction.Rollback()
// 		}
// 	}()

// 	if !ok || err != nil {
// 		if err == nil {
// 			err = errors.New("something went wrong on delete task")
// 		}
// 		transaction.Rollback()
// 		return nil, utils.GenerateError(ts.Name, err.Error())
// 	}

// 	newTask, err := repository.Create[entities.Task](repository.Tasks, entities.Task{
// 		EventId:   currentEvent.ID,
// 		AccountId: currentAcc.ID,
// 		Type:      currentTask.Type,
// 		Provider:  currentTask.Provider,
// 		Date:      currentTask.Date.AddDate(1, 0, 0),
// 	}, transaction)

// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, utils.GenerateError(ts.Name, err.Error())
// 	}

// 	log.Println("task creted from cron" + newTask.ID.String())

// 	textMsg := "Attention!" + " For " + currentTask.Type + " in " + currentEvent.Date.Format("01-02") + " will be event " + currentEvent.Info

// 	transaction.Commit()

// 	return &InfoAboutTaskForTgProvider{
// 		ChatId: chatId,
// 		Text:   textMsg,
// 	}, nil

// }
