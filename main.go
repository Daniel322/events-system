package main

import (
	"events-system/internal/controllers"
	"events-system/internal/domain"
	db "events-system/internal/providers/db"
	"events-system/internal/providers/server"
	"events-system/internal/providers/telegram"
	"events-system/internal/repositories"
	"events-system/internal/services"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// init external providers
	db := db.NewDatabase(os.Getenv("GOOSE_DBSTRING"))
	server := server.NewEchoInstance()

	// init domain factories
	userFactory := domain.NewUserFactory()
	accountFactory := domain.NewAccountFactory()
	eventFactory := domain.NewEventFactory()
	taskFactory := domain.NewTaskFactory()

	// init repositories
	userRepository := repositories.NewRepository("UserRepository", db, userFactory)
	accountRepository := repositories.NewRepository("AccountRepository", db, accountFactory)
	eventRepository := repositories.NewRepository("EventRepository", db, eventFactory)
	taskRepository := repositories.NewRepository("TaskRepository", db, taskFactory)

	// init services
	userService := services.NewUserService(db, userRepository, accountRepository)
	accountService := services.NewAccountService(db, accountRepository)
	eventsService := services.NewEventService(db, eventRepository, taskRepository)
	userController := controllers.NewUserController(
		server.Instance,
		userService,
	)

	tgBotProvider, err := telegram.NewTgBotProvider(os.Getenv("TG_BOT_TOKEN"), userService, accountService, eventsService)

	if err != nil {
		panic(err.Error())
	}

	tgBotProvider.Bootstrap()

	// init http routes
	userController.InitRoutes()

	// start http server
	server.Start(os.Getenv("HTTP_PORT"))
}
