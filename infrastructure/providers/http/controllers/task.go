package controllers

import (
	"events-system/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	Name    string
	server  *echo.Echo
	service interfaces.InternalUsecase
}

func NewTaskController(server *echo.Echo, service interfaces.InternalUsecase) *TaskController {
	return &TaskController{
		server:  server,
		Name:    "TaskController",
		service: service,
	}
}

func (controller *TaskController) GetTodayTasks(c echo.Context) error {
	tasks, err := controller.service.GetListOfTodayTasks()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func (controller *TaskController) ExecTask(c echo.Context) error {
	return c.JSON(http.StatusCreated, "ok")
}

func (controller *TaskController) InitRoutes() {
	controller.server.GET("/tasks/today", controller.GetTodayTasks)
	controller.server.POST("/tasks/:id", controller.ExecTask)
}
