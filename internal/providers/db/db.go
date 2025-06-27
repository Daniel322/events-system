package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseInstance = *gorm.DB

type Database struct {
	// TODO: change gorm to special astract interface
	Instance DatabaseInstance
	Url      string
}

func NewDatabase(url string) *Database {
	conn, err := gorm.Open(postgres.Open(url), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	return &Database{
		Url:      url,
		Instance: conn,
	}
}

func (db *Database) CreateTransaction() DatabaseInstance {
	tx := db.Instance.Begin()

	return tx
}
