package main

import (
	"encoding/json"
	"events-system/modules/db"
	event_module "events-system/modules/event"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
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

	// telegram_api.BootstrapBot()

	uuid, err := uuid.Parse("92e7e817-275a-4fe5-bf59-da72641c8549")

	if err != nil {
		log.Fatal(err)
	}

	result, err := event_module.GetUserEvents(&uuid)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	firstEvent := (*result)[0]

	var jsonNotifyLevels []string

	json.Unmarshal(firstEvent.NotifyLevels, &jsonNotifyLevels)

	fmt.Println(jsonNotifyLevels)
}
