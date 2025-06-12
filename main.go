package main

import (
	db "events-system/internal/providers/db"
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

	db := db.NewDatabase(os.Getenv("GOOSE_DBSTRING"))

	fmt.Println(db)

	// telegram_api.BootstrapBot()

	server := server.NewEchoInstance()

	server.Start(os.Getenv("HTTP_PORT"))
}
