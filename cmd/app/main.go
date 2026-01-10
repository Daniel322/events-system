package main

import (
	"context"
	pg_db "events-system/infrastructure/providers/db/postgres"
	server "events-system/infrastructure/providers/http"
	"events-system/internal/components"
	"events-system/pkg/config"
	"os"
)

// TODO: repository abstraction
// TODO: components slice
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

	userFactory := components.NewUserFactory(db_adapter)

	user := userFactory.NewUser("test factoryzzzz")

	ctx := context.Background()

	user.Save(ctx)

	user.Username = "asdcdcd"

	user.Save(ctx)

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
