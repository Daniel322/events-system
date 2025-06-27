package main

import (
	"events-system/internal/controllers"
	"events-system/internal/domain"
	db "events-system/internal/providers/db"
	"events-system/internal/providers/server"
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

	// telegram_api.BootstrapBot()

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

	fmt.Println(eventRepository, taskRepository)

	userService := services.NewUserService(db, userRepository, accountRepository)
	userController := controllers.NewUserController(
		server.Instance,
		userService,
	)

	// init http routes
	userController.InitRoutes()

	// start http server
	server.Start(os.Getenv("HTTP_PORT"))
}
