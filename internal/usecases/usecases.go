package usecases

import (
	"events-system/interfaces"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"log"
	"strconv"
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
	"tommorow": time.Hour * 24,
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

func (usecase *InternalUseCases) CreateUser(data dto.CreateUserInput) (*dto.OutputUser, error) {
	transaction := usecase.BaseRepository.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	user, err := usecase.UserService.Create(data.Username, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError("CreateUser", err.Error())
	}

	acc, err := usecase.AccountService.Create(dto.CreateAccountData{
		UserId:    user.ID,
		AccountId: data.AccountId,
		Type:      data.Type,
	}, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError("CreateUser", err.Error())
	}

	var accs []entities.Account
	accs = append(accs, *acc)

	if trRes := transaction.Commit(); trRes.Error != nil {
		return nil, utils.GenerateError("CreateUser", trRes.Error.Error())
	}

	return &dto.OutputUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  accs,
	}, nil
}

func (usecase *InternalUseCases) GetUser(id string) (*dto.OutputUser, error) {
	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	users, err := usecase.UserService.Find(findOptions)

	if err != nil {
		return nil, utils.GenerateError("GetUser", err.Error())
	}

	if len(*users) == 0 {
		return nil, utils.GenerateError("GetUser", "user not found")
	}

	user := (*users)[0]

	options := make(map[string]interface{})

	options["user_id"] = user.ID

	accs, err := usecase.AccountService.Find(options)

	if err != nil {
		return nil, utils.GenerateError("GetUser", err.Error())
	}

	return &dto.OutputUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  *accs,
	}, nil
}

func (usecase *InternalUseCases) CheckTGAccount(accountId int64) (*entities.Account, error) {
	options := make(map[string]interface{})
	options["account_id"] = strconv.Itoa(int(accountId))
	currentAccounts, err := usecase.AccountService.Find(options)

	if err != nil {
		return nil, utils.GenerateError("CheckTGAccount", err.Error())
	}

	if len(*currentAccounts) == 0 {
		return nil, nil
	}

	return &(*currentAccounts)[0], nil
}

func (usecase *InternalUseCases) GetListOfTodayTasks() (*[]entities.Task, error) {
	options := make(map[string]interface{})
	options["date"] = time.Now().Format("2006-01-02")

	tasks, err := usecase.TaskService.Find(options)

	if err != nil {
		return nil, utils.GenerateError("GetListOfTodayTasks", err.Error())
	}

	return tasks, nil
}

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

	// unreal error case
	notifyLevels, _ := event.NotifyLevels.Value()

	// unreal error case
	providers, _ := event.Providers.Value()

	return &dto.OutputEvent{
		ID:           event.ID,
		UserId:       event.UserId,
		Info:         event.Info,
		Date:         event.Date,
		NotifyLevels: notifyLevels.(string),
		Providers:    providers.(string),
		CreatedAt:    event.CreatedAt,
		UpdatedAt:    event.UpdatedAt,
		Tasks:        tasks,
	}, nil
}

func (usecase *InternalUseCases) ExecTask(taskId string) (*dto.InfoAboutTaskForTgProvider, error) {
	taskFindOptions := make(map[string]interface{})
	taskFindOptions["id"] = taskId

	currentTask, err := usecase.TaskService.FindOne(taskFindOptions)

	if err != nil {
		return nil, utils.GenerateError("ExecTask", err.Error())
	}

	eventFindOptions := make(map[string]interface{})
	eventFindOptions["id"] = currentTask.EventId.String()

	currentEvent, err := usecase.EventService.FindOne(eventFindOptions)

	if err != nil {
		return nil, utils.GenerateError("ExecTask", err.Error())
	}

	accFindOptions := make(map[string]interface{})
	accFindOptions["id"] = currentTask.AccountId.String()

	currentAcc, err := usecase.AccountService.FindOne(accFindOptions)

	if err != nil {
		return nil, utils.GenerateError("ExecTask", err.Error())
	}

	transaction := usecase.BaseRepository.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	ok, err := usecase.TaskService.Delete(currentTask.ID.String(), transaction)

	if !ok || err != nil {
		transaction.Rollback()
		if err != nil {
			return nil, utils.GenerateError("ExecTask", err.Error())
		}
		return nil, utils.GenerateError("ExecTask", "task not deleted")
	}

	newTask, err := usecase.TaskService.Create(
		dto.CreateTaskData{
			EventId:   currentEvent.ID,
			AccountId: currentAcc.ID,
			Type:      currentTask.Type,
			Provider:  currentTask.Provider,
			Date:      currentTask.Date.AddDate(1, 0, 0),
		},
		transaction,
	)

	if err != nil {
		return nil, utils.GenerateError("ExecTask", err.Error())
	}

	log.Println("task creted from cron" + newTask.ID.String())

	textMsg := "Attention!" + " For " + currentTask.Type + " in " + currentEvent.Date.Format("01-02") + " will be event " + currentEvent.Info

	transaction.Commit()

	return &dto.InfoAboutTaskForTgProvider{
		ChatId: currentAcc.AccountId,
		Text:   textMsg,
	}, nil
}
