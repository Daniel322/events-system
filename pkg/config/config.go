package config

import (
	"events-system/pkg/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Name   string
	Logger *log.Logger
}

func Bootstrap() (*Config, error) {
	logger := log.New(os.Stdout, "Config ", log.LstdFlags)

	config := Config{
		Name:   "Config",
		Logger: logger,
	}

	err := godotenv.Load()

	if err != nil {
		config.Logger.Println("Error loading .env file", err)

		return nil, err
	}

	return &config, nil
}

func (cfg *Config) get(field string) (string, error) {
	field_value := os.Getenv(field)

	if len(field_value) == 0 {
		return field_value, utils.GenerateError(cfg.Name, "error on get "+field)
	}

	return field_value, nil
}

func (cfg *Config) DB_URL() (string, error) {
	return cfg.get(DB_URL)
}

func (cfg *Config) TG_TOKEN() (string, error) {
	return cfg.get(TG_TOKEN)
}

func (cfg *Config) HTTP_PORT() (string, error) {
	return cfg.get(HTTP_PORT)
}

func (cfg *Config) CRON_INTERVAL() (string, error) {
	return cfg.get(CRON_INTERVAL)
}
