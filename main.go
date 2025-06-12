package main

import (
	db "events-system/internal/providers"
	"events-system/internal/providers/server"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db.ConnectDatabase()

	// telegram_api.BootstrapBot()

	server := server.NewEchoInstance()

	server.Start(os.Getenv("HTTP_PORT"))
}
