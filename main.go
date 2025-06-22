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
	// init user domain
	userFactory := domain.NewUserFactory("user-factory")
	userRepository := repositories.NewUserRepository("userRepository", db.Instance, userFactory)
	fmt.Println(userRepository)
	userService := services.NewUserService("users", userRepository)
	userController := controllers.NewUserController(
		server.Instance,
		userService,
	)

	// init http routes
	userController.InitRoutes()

	fmt.Println(server.Instance.Routes())

	// start http server
	server.Start(os.Getenv("HTTP_PORT"))
}
