package main

import (
	"events-system/infrastructure/providers/http/server"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// TODO: logger provider
// TODO: config manager
// TODO: repository abstraction
// TODO: components slice
// TODO: graceful shutdown

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// db_url := os.Getenv("GOOSE_DBSTRING")

	// conn, err := gorm.Open(postgres.Open(db_url), &gorm.Config{
	// 	SkipDefaultTransaction: true,
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected!")

	// // init external providers
	// database_instance := db.NewDatabase(db_url, conn)
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
