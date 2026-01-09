package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Bootstrap() error {
	logger := log.New(os.Stdout, "Config ", log.LstdFlags)
	err := godotenv.Load()
	if err != nil {
		logger.Println("Error loading .env file", err)
	}

	return err
}
