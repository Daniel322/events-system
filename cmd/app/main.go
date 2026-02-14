package main

import (
	"context"
	pg_db "events-system/infrastructure/providers/db/postgres"
	server "events-system/infrastructure/providers/http"
	"events-system/internal/application"
	"events-system/internal/domain/account"
	"events-system/internal/domain/user"
	"events-system/pkg/config"
	"events-system/pkg/vo"
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
	// eventRepo := event.NewEventsRepo(db_adapter)
	// taskRepo := task.NewTaskRepo(db_adapter)

	createUserAction := application.NewCreateUser(userRepo, accRepo)

	ctx := context.Background()

	userName, _ := vo.NewNonEmptyString("Daniil")
	typ, _ := account.NewAccountType("mail")
	val, _ := account.NewAccountValue("kravchenkodanil122@gmail.com", typ)
	user, err := createUserAction.Run(ctx, application.CreateUserState{
		Username:     userName,
		Type:         typ,
		AccountValue: val,
	},
	)

	fmt.Println(string(user.ToJSON()))

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
