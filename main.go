package main

import (
	"context"
	"events-system/modules/db"
	user_module "events-system/modules/user"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var err error

func main() {
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Hello!")
	fmt.Println("try to connect to db")

	db.ConnectDatabase(context.Background())

	defer db.Close(context.Background())

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

	// test user.CreateUser

	var testResult *user_module.User

	testResult, err = user_module.CreateUser(user_module.CreateUserData{Username: "dkravchenkoo"})

	fmt.Println(testResult)
}
