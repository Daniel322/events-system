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
	userFactory := domain.NewUserFactory("user-factory")
	accountFactory := domain.NewAccountFactory()
	fmt.Println(accountFactory)
	// init user domain

	userRepository := repositories.NewUserRepository("userRepository", db.Instance, userFactory)
	userService := services.NewUserService("users", userRepository)
	userController := controllers.NewUserController(
		server.Instance,
		userService,
	)

	user, err := userRepository.UpdateUser("53e6ea62-3ea3-453f-91d8-21a1ed7b4381", domain.UserData{Username: "asdasda"})

	fmt.Println(user, err)

	// init http routes
	userController.InitRoutes()

	// start http server
	server.Start(os.Getenv("HTTP_PORT"))
}
