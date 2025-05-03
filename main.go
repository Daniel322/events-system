package main

import (
	"context"
	telegram_api "events-system/apis/telegram"
	"events-system/modules/db"
	event_module "events-system/modules/event"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
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
	fmt.Println("Hello!")
	fmt.Println("try to connect to db")

	db.ConnectDatabase(context.Background())

	defer db.Close(context.Background())

	telegram_api.BootstrapBot()

	var rows pgx.Rows
	var result []string

	if db.CheckConnection() {
		fmt.Println("Connected!")
	} else {
		fmt.Println("Disconnected!")
	}

	rows, err = db.Connection.Query(context.Background(), "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';")
	if err != nil {
		log.Fatal(err)
	}

	var startLoop = false
	for !startLoop {
		needNext := rows.Next()
		if needNext {
			var rowResult []any
			rowResult, err = rows.Values()
			if err != nil {
				log.Fatal(err)
			}
			rowResultEl := rowResult[0]
			result = append(result, rowResultEl.(string))
			fmt.Println(rowResultEl)
		} else {
			startLoop = true
		}
	}

	fmt.Println(result)

	event_module.CreateEvent(event_module.CreateEventData{
		UserId:       "92e7e817-275a-4fe5-bf59-da72641c8549",
		Info:         "Moms birthdays",
		Date:         "1975-08-28T00:00:00.000Z",
		NotifyLevels: []string{"month", "week", "tomorrow", "today"},
		Providers:    []string{"telegram"},
	}, context.Background())
}
