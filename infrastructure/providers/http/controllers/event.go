package controllers

import (
	"events-system/interfaces"
	"events-system/internal/dto"
	"events-system/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventController struct {
	Name    string
	server  *echo.Echo
	service interfaces.InternalUsecase
}

func NewEventController(server *echo.Echo, service interfaces.InternalUsecase) *EventController {
	return &EventController{
		server:  server,
		Name:    "EventController",
		service: service,
	}
}

func (controller *EventController) CreateEvent(c echo.Context) error {
	eventData := new(dto.CreateEventDTO)

	err := c.Bind(eventData)
	if err != nil {
		generatedError := utils.GenerateError(controller.Name, err.Error())
		return c.String(http.StatusBadRequest, generatedError.Error())
	}

	err = c.Validate(eventData)
	if err != nil {
		generatedError := utils.GenerateError(controller.Name, err.Error())
		return c.String(http.StatusBadRequest, generatedError.Error())
	}

	event, err := controller.service.CreateEvent(*eventData)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, event)
}

// TODO: add get and patch methods in s-2.0
func (controller *EventController) InitRoutes() {
	controller.server.POST("/events", controller.CreateEvent)
}
