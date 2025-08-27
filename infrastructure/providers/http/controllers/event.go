package controllers

import (
	"events-system/interfaces"
	"events-system/internal/dto"
	dependency_container "events-system/pkg/di"
	"events-system/pkg/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventController struct {
	Name   string
	server *echo.Echo
}

func NewEventController(server *echo.Echo) *EventController {
	controller := &EventController{
		server: server,
		Name:   "EventController",
	}

	dependency_container.Container.Add("eventController", controller)

	return controller
}

func (con *EventController) ExecRoute(c echo.Context) error {
	eventService, err := dependency_container.Container.Get("eventService")

	if err != nil {
		err = utils.GenerateError(con.Name, err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}
	switch method := c.Request().Method; method {
	case "GET":
		fmt.Println("GET method not implemented")
	case "POST":
		eventData := new(dto.CreateEventDTO)

		err := c.Bind(eventData)
		if err != nil {
			generatedError := utils.GenerateError(con.Name, err.Error())
			return c.String(http.StatusBadRequest, generatedError.Error())
		}

		err = c.Validate(eventData)
		if err != nil {
			generatedError := utils.GenerateError(con.Name, err.Error())
			return c.String(http.StatusBadRequest, generatedError.Error())
		}

		event, err := eventService.(interfaces.EventService).CreateEvent(*eventData)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, event)
	case "PATCH":
		fmt.Println("PATCH method not implemented")
	case "PUT":
		fmt.Println("PUT method not implemented")
	case "DELETE":
		fmt.Println("DELETE method not implemented")
	}
	return c.JSON(200, "ok")
}

func (con *EventController) InitRoutes() {
	con.server.POST("/event", con.ExecRoute)
}
