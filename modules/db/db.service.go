package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var err error
var Connection *pgx.Conn

func ConnectDatabase(context context.Context) {
	Connection, err = pgx.Connect(context, os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatal(err)
	}
}

func Close(context context.Context) {
	Connection.Close(context)
}
