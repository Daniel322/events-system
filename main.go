package main

import (
	"events-system/internal/interfaces/http/controllers"
	db "events-system/internal/providers/db"
	"events-system/internal/providers/server"
	"events-system/internal/services"
	usecases "events-system/internal/usecase"
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
	userService := services.NewUserService()
	userUseCase := usecases.NewUserUseCase(
		db.Instance,
		userService,
	)
	userController := controllers.NewUserController(
		server.Instance,
		userUseCase,
	)

	// init http routes
	userController.InitRoutes()

	fmt.Println(server.Instance.Routes())

	// start http server
	server.Start(os.Getenv("HTTP_PORT"))
}
