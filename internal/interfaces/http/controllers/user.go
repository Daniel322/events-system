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
