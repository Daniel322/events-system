package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error
var Connection *gorm.DB

func ConnectDatabase() {
	Connection, err = gorm.Open(postgres.Open(os.Getenv("GOOSE_DBSTRING")))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
}
