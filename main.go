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

	// init http routes
	userController.InitRoutes()

	account, err := accountFactory.Create(domain.CreateAccountData{
		UserId:    "0bf81a5a-6b50-4b04-a6d9-5828d3ca9b72",
		AccountId: "sdafadsfas",
		Type:      "http",
	})

	fmt.Println(account, err)

	// start http server
	server.Start(os.Getenv("HTTP_PORT"))
}
