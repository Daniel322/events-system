package main

import (
	db "events-system/internal/providers"
	"fmt"
	"os"

	"net/http"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db.ConnectDatabase()

	// telegram_api.BootstrapBot()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("HTTP_PORT")))

}
