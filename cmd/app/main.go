package main

import (
	"context"
	pg_db "events-system/infrastructure/providers/db/postgres"
	server "events-system/infrastructure/providers/http"
	"events-system/internal/application/commands"
	"events-system/internal/application/queries"
	"events-system/internal/domain/account"
	"events-system/internal/domain/event"
	"events-system/internal/domain/user"
	"events-system/pkg/config"
	"fmt"
	"os"
)

// TODO: graceful shutdown

func main() {

	err := config.Config.Bootstrap()

	if err != nil {
		panic(err.Error())
	}

	db_conn, err := pg_db.Connect()

	if err != nil {
		panic(err.Error())
	}

	db_adapter := pg_db.NewDbAdapter(db_conn)

	userRepo := user.NewUsersRepo(db_adapter)
	accRepo := account.NewAccRepo(db_adapter)
	eventRepo := event.NewEventsRepo(db_adapter)
	// taskRepo := task.NewTaskRepo(db_adapter)

	createUserAction := commands.NewCreateUser(userRepo, accRepo)
	getUserAction := queries.NewGetUser(userRepo, accRepo, eventRepo)

	ctx := context.Background()

	state, _ := createUserAction.Validate(commands.CreateUserData{
		Username:     "Daniil",
		Type:         "mail",
		AccountValue: "kravchenkodanil12342@gmail.com",
	})

	user, err := createUserAction.Run(ctx, *state)

	// fmt.Println(string(user.ToJSON()))

	userG, err := getUserAction.Run(ctx, user.ID.String())

	fmt.Println(userG, string(userG.ToJSON()))

	server := server.NewEchoInstance()

	// // init di container

	// dependency_container := dependency_container.NewDIContainer()

	// internalUseCases := initInternalDependencies(dependency_container, database_instance)

	// // init controllers
	// userController := controllers.NewUserController(
	// 	server.Instance,
	// 	internalUseCases,
	// )
	// eventController := controllers.NewEventController(server.Instance, internalUseCases)
	// taskController := controllers.NewTaskController(server.Instance, internalUseCases)

	// tgBotProvider, err := telegram.NewTgBotProvider(os.Getenv("TG_BOT_TOKEN"), internalUseCases)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// cronProvider := cron.NewCronProvider(tgBotProvider, internalUseCases)

	// cronProvider.Bootstrap()

	// go tgBotProvider.Bootstrap()

	// // init http routes
	// userController.InitRoutes()
	// eventController.InitRoutes()
	// taskController.InitRoutes()

	// start http server
	server.Start(os.Getenv("HTTP_PORT"))

	// <-ctx.Done()

	// log.Println("shutting down server gracefully")

	// shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// database_instance.Close()
	// tgBotProvider.Close()
	// server.Close(shutdownCtx)
}
