package controllers

import (
	usecases "events-system/internal/usecase"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	server  *echo.Echo
	useCase *usecases.UserUseCase
}

func NewUserController(server *echo.Echo, useCase *usecases.UserUseCase) *UserController {
	return &UserController{
		server:  server,
		useCase: useCase,
	}
}

func (uc UserController) ExecRoute(c echo.Context) error {
	return c.JSON(200, "ok")
}

func (uc UserController) InitRoutes() {
	uc.server.POST("/users", uc.ExecRoute)
	uc.server.GET("/users/:id", uc.ExecRoute)
	uc.server.PATCH("/users/:id", uc.ExecRoute)
	uc.server.DELETE("/users/:id", uc.ExecRoute)
}
