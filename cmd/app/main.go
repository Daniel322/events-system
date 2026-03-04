package main

import (
	"context"
	"events-system/infrastructure/cache"
	"events-system/infrastructure/config"
	"events-system/infrastructure/cron"
	pg_db "events-system/infrastructure/db/adapters/postgres"
	"events-system/infrastructure/telegram"
	"events-system/internal/application/commands"
	"events-system/internal/application/queries"
	"events-system/internal/domain/account"
	"events-system/internal/domain/event"
	"events-system/internal/domain/task"
	"events-system/internal/domain/user"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := config.Config.Bootstrap()

	if err != nil {
		panic(err.Error())
	}

	db_conn, err := pg_db.Connect()

	if err != nil {
		panic(err.Error())
	}

	pg_db.InitAdapter(db_conn)

	cache.Init()

	user.InitRepo(pg_db.Adapter)
	account.InitRepo(pg_db.Adapter)
	event.InitRepo(pg_db.Adapter)
	task.InitRepo(pg_db.Adapter)

	commands.InitCreateUser()
	commands.InitCreateEvent()
	commands.InitExecTask()
	queries.InitGetUser()
	queries.InitTasksList()
	queries.InitCheckAccount()
	queries.InitEventsList()

	err = telegram.NewTgBotProvider()

	if err != nil {
		panic(err.Error())
	}

	cronProvider := cron.NewCronProvider(telegram.Provider)

	cronProvider.Bootstrap()

	go telegram.Provider.Bootstrap()

	// state, _ := createEventAction.Validate(commands.CreateEventData{
	// 	UserId:       "9bd6d11f-c4b2-4863-93ab-09dbd7728880",
	// 	AccId:        "0f6af9d8-af28-42ee-b895-417901cd70a1",
	// 	Info:         "app test event with tasks",
	// 	Date:         time.Now(),
	// 	NotifyLevels: []string{"today", "tomorrow"},
	// 	Providers:    []string{"telegram"},
	// })

	// event, err := createEventAction.Run(ctx, state)

	// fmt.Println(event.ToPlain())

	// state, _ := createUserAction.Validate(commands.CreateUserData{
	// 	Username:     "Daniil",
	// 	Type:         "mail",
	// 	AccountValue: "kravchenkodanil12342@gmail.com",
	// })

	// user, err := createUserAction.Run(ctx, *state)

	// userG, err := getUserAction.Run(ctx, user.ID.String())

	// fmt.Println(userG)

	<-ctx.Done()

	log.Println("shutting down server gracefully")

	cronProvider.Stop()
	telegram.Provider.Close()
	pg_db.Close(db_conn)
}
