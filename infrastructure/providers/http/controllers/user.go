package controllers

import (
	"events-system/interfaces"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Name    string
	server  *echo.Echo
	service interfaces.InternalUsecase
}

func NewUserController(server *echo.Echo, service interfaces.InternalUsecase) *UserController {
	return &UserController{
		Name:    "UserController",
		server:  server,
		service: service,
	}
}

func (controller UserController) CreateUser(c echo.Context) error {
	userData := new(dto.UserDataDTO)
	err := c.Bind(userData)
	if err != nil || len(userData.Username) == 0 {
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = c.Validate(userData)
	if err != nil || len(userData.Username) == 0 {
		generatedError := utils.GenerateError(controller.Name, err.Error())
		return c.String(http.StatusBadRequest, generatedError.Error())
	}

	user, err := controller.service.CreateUser(dto.CreateUserInput{
		Username:  userData.Username,
		AccountId: userData.AccountId,
		Type:      entities.AccountType(0),
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (controller UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := controller.service.GetUser(id)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (controller UserController) InitRoutes() {
	controller.server.POST("/users", controller.CreateUser)
	controller.server.GET("/users/:id", controller.GetUser)
}
