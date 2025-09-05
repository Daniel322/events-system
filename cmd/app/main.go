package main

import (
	"context"
	db "events-system/infrastructure/providers/db"
	"events-system/infrastructure/providers/http/server"
	entities "events-system/internal/entity"
	"events-system/internal/services"
	dependency_container "events-system/pkg/di"
	"events-system/pkg/repository"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func initDependencies(di *dependency_container.DependencyContainer, db *db.Database) {

	// init base repository
	base_repository := repository.NewBaseRepository(db)

	// init repos
	user_repository := repository.NewRepository[entities.User](repository.Users, base_repository)
	account_repository := repository.NewRepository[entities.Account](repository.Accounts, base_repository)
	event_repository := repository.NewRepository[entities.Event](repository.Events, base_repository)
	task_repository := repository.NewRepository[entities.Task](repository.Tasks, base_repository)

	user_service := services.NewUserService(user_repository)
	account_service := services.NewAccountService(account_repository)
	event_service := services.NewEventService(event_repository)
	task_service := services.NewTaskService(task_repository)

	di.Add(
		"baseRepository",
		base_repository,
	)

	di.Add(
		"userService",
		user_service,
	)

	di.Add(
		"accountService",
		account_service,
	)

	di.Add(
		"eventService",
		event_service,
	)

	di.Add(
		"taskService",
		task_service,
	)

}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer stop()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// init external providers
	database_instance := db.NewDatabase(os.Getenv("GOOSE_DBSTRING"))
	server := server.NewEchoInstance()

	// init di container

	dependency_container := dependency_container.NewDIContainer()

	initDependencies(dependency_container, database_instance)

	// init domain factories
	// entities.NewAccountFactory()
	// entities.NewEventFactory()
	// entities.NewTaskFactory()

	// init services v2
	// services.NewUserService()

	// init services
	// userService := services.NewUserService()
	// accountService := services.NewAccountService()
	// eventsService := services.NewEventService()
	// tasksService := services.NewTaskService()

	// init controllers
	// userController := controllers.NewUserController(
	// 	server.Instance,
	// 	userService,
	// )
	// eventController := controllers.NewEventController(server.Instance)

	// tgBotProvider, err := telegram.NewTgBotProvider(os.Getenv("TG_BOT_TOKEN"), userService, accountService, eventsService)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// cronProvider := cron.NewCronProvider(tgBotProvider, tasksService)

	// cronProvider.Bootstrap()

	// go tgBotProvider.Bootstrap()

	// // init http routes
	// userController.InitRoutes()
	// eventController.InitRoutes()

	// start http server
	go server.Start(os.Getenv("HTTP_PORT"))

	<-ctx.Done()

	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	database_instance.Close()
	// tgBotProvider.Close()
	server.Close(shutdownCtx)
}
