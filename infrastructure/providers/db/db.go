package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseInstance = *gorm.DB

type Database struct {
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

func (db *Database) Close() error {
	dbInstance, err := db.Instance.DB()

	if err != nil {
		return err
	}

	err = dbInstance.Close()

	if err == nil {
		log.SetPrefix("DB ")
		log.Println("Close DB connection")
	}

	return err
}
