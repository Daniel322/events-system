package main

import (
	server "events-system/infrastructure/providers/http"
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
	username, _ := vo.NewNonEmptyString("test")
	user := user.New(username)

	fmt.Println(user.Username(), user.ID)

	// db_conn, err := pg_db.Connect()

	// if err != nil {
	// 	panic(err.Error())
	// }

	// db_adapter := pg_db.NewDbAdapter(db_conn)

	// users := components.NewUsersFactory(db_adapter)

	// user := users.NewUser("zxccxz")

	// tx := db_adapter.Instance.Begin()

	// ctx := context.WithValue(context.Background(), "transaction", tx)

	// user.Save(ctx)

	// user.Username = "asdcdcd"

	// user.Save(ctx)

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
