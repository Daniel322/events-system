package main

import (
	telegram_api "events-system/apis/telegram"
	"events-system/modules/db"
	"fmt"
	"sync"

	"github.com/joho/godotenv"
)

var err error

func main() {
	var mutex = &sync.RWMutex{}
	mutex.Lock()

	defer mutex.Unlock()
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db.ConnectDatabase()

	telegram_api.BootstrapBot()
}
