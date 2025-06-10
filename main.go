package main

import (
	"events-system/modules/db"
	"fmt"
	"sync"

	"net/http"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

var err error

func main() {
	var mutex = &sync.RWMutex{}
	mutex.Lock()

	defer mutex.Unlock()
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db.ConnectDatabase()

	// telegram_api.BootstrapBot()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
