package config

import (
	"events-system/pkg/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IConfig struct {
	Name   string
	Logger *log.Logger
}

var logger = log.New(os.Stdout, "Config ", log.LstdFlags)

var Config = IConfig{
	Name:   "Config",
	Logger: logger,
}

func (cfg *IConfig) Bootstrap() error {

	err := godotenv.Load()

	if err != nil {
		Config.Logger.Println("Error loading .env file", err)

		return err
	}

	return nil
}

func (cfg *IConfig) get(field string) (string, error) {
	field_value := os.Getenv(field)

	if len(field_value) == 0 {
		return field_value, utils.GenerateError(cfg.Name, "error on get "+field)
	}

	return field_value, nil
}

func (cfg *IConfig) DB_URL() (string, error) {
	return cfg.get(DB_URL)
}

func (cfg *IConfig) TG_TOKEN() (string, error) {
	return cfg.get(TG_TOKEN)
}

func (cfg *IConfig) HTTP_PORT() (string, error) {
	return cfg.get(HTTP_PORT)
}

func (cfg *IConfig) CRON_INTERVAL() (string, error) {
	return cfg.get(CRON_INTERVAL)
}

func (cfg *IConfig) EXAMPLE_FILE_PATH() (string, error) {
	return cfg.get(EXAMPLE_FILE_PATH)
}
