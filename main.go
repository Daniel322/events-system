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

	db := db.NewDatabase(os.Getenv("GOOSE_DBSTRING"))

	// telegram_api.BootstrapBot()

	// init http server
	server := server.NewEchoInstance()
	// init domain factories
	userFactory := domain.NewUserFactory()
	accountFactory := domain.NewAccountFactory()
	eventFactory := domain.NewEventFactory()
	taskFactory := domain.NewTaskFactory()
	fmt.Println(accountFactory, eventFactory, taskFactory)

	// init repositories
	userRepository := repositories.NewUserRepository(db, userFactory)

	userService := services.NewUserService("users", userRepository)
	userController := controllers.NewUserController(
		server.Instance,
		userService,
	)

	// init http routes
	userController.InitRoutes()

	// start http server
	server.Start(os.Getenv("HTTP_PORT"))
}
