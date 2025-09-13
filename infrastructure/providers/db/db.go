package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type DatabaseInstance = *gorm.DB

type Database struct {
	Instance DatabaseInstance
	Url      string
}

func NewDatabase(url string, connection DatabaseInstance) *Database {
	fmt.Println("Connected!")

	return &Database{
		Url:      url,
		Instance: connection,
	}
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
