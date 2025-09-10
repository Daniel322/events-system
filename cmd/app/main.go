package main

import (
	"context"
	"events-system/infrastructure/providers/cron"
	db "events-system/infrastructure/providers/db"
	"events-system/infrastructure/providers/http/controllers"
	"events-system/infrastructure/providers/http/server"
	"events-system/infrastructure/providers/telegram"
	"events-system/interfaces"
	entities "events-system/internal/entity"
	"events-system/internal/services"
	"events-system/internal/usecases"
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

func initInternalDependencies(
	di *dependency_container.DependencyContainer,
	db *db.Database,
) interfaces.InternalUsecase {

	// init base repository
	base_repository := repository.NewBaseRepository(db)

	// init repos
	user_repository := repository.NewRepository[entities.User](repository.Users, base_repository)
	account_repository := repository.NewRepository[entities.Account](repository.Accounts, base_repository)
	event_repository := repository.NewRepository[entities.Event](repository.Events, base_repository)
	task_repository := repository.NewRepository[entities.Task](repository.Tasks, base_repository)

	// init services
	user_service := services.NewUserService(user_repository)
	account_service := services.NewAccountService(account_repository)
	event_service := services.NewEventService(event_repository)
	task_service := services.NewTaskService(task_repository)

	internalUseCases := usecases.NewInternalUseCases(
		base_repository,
		user_service,
		account_service,
		event_service,
		task_service,
	)

	di.Add("useCases", internalUseCases)

	return internalUseCases
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

	internalUseCases := initInternalDependencies(dependency_container, database_instance)

	// init controllers
	userController := controllers.NewUserController(
		server.Instance,
		internalUseCases,
	)
	eventController := controllers.NewEventController(server.Instance, internalUseCases)
	taskController := controllers.NewTaskController(server.Instance, internalUseCases)

	tgBotProvider, err := telegram.NewTgBotProvider(os.Getenv("TG_BOT_TOKEN"), internalUseCases)

	if err != nil {
		panic(err.Error())
	}

	cronProvider := cron.NewCronProvider(tgBotProvider, internalUseCases)

	cronProvider.Bootstrap()

	go tgBotProvider.Bootstrap()

	// init http routes
	userController.InitRoutes()
	eventController.InitRoutes()
	taskController.InitRoutes()

	// start http server
	go server.Start(os.Getenv("HTTP_PORT"))

	<-ctx.Done()

	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	database_instance.Close()
	tgBotProvider.Close()
	server.Close(shutdownCtx)
}
