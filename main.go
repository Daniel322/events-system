package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var err error
var conn *pgx.Conn

func main() {
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Hello!")
	fmt.Println("try to connect to db")

	conn, err = pgx.Connect(context.Background(), os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	var rows pgx.Rows
	var result []string

	rows, err = conn.Query(context.Background(), "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';")
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

	// result, err = rows.Values()

	fmt.Println(result)

	fmt.Println("Connected!")
}
