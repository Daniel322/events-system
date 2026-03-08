package pg_db

import (
	"events-system/infrastructure/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db_url, err := config.Config.DB_URL()

	if err != nil {
		return nil, err
	}

	conn, err := gorm.Open(postgres.Open(db_url), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, err
	}

	log.SetPrefix("DB Adapter ")
	log.Println("Connected to Database!")

	return conn, nil
}

func Close(db *gorm.DB) error {
	instance, err := db.DB()

	if err != nil {
		return err
	}

	err = instance.Close()

	if err == nil {
		log.SetPrefix("DB Adapter ")
		log.Println("Close DB connection")
	}

	return err
}
