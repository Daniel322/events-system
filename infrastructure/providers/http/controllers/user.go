package controllers

import (
	"events-system/internal/dto"
	"events-system/internal/interfaces"
	"events-system/internal/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Name        string
	server      *echo.Echo
	userService interfaces.IUserService
}

func NewUserController(server *echo.Echo, service interfaces.IUserService) *UserController {
	return &UserController{
		Name:        "UserController",
		server:      server,
		userService: service,
	}
}

func (uc UserController) ExecRoute(c echo.Context) error {
	fmt.Println(c.Request().Method, c.Path())
	switch method := c.Request().Method; method {
	case "GET":
		fmt.Println("GET METHOD")
		id := c.Param("id")
		fmt.Println(id)
		user, err := uc.userService.GetUser(id)

		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.JSON(200, user)
	case "POST":
		fmt.Println("start post method")
		userData := new(dto.UserDataDTO)
		err := c.Bind(userData)
		if err != nil || len(userData.Username) == 0 {
			return c.String(http.StatusBadRequest, "bad request")
		}

		err = c.Validate(userData)
		if err != nil || len(userData.Username) == 0 {
			generatedError := utils.GenerateError(uc.Name, err.Error())
			return c.String(http.StatusBadRequest, generatedError.Error())
		}

		user, err := uc.userService.CreateUser(dto.UserDataDTO{
			Username:  userData.Username,
			AccountId: userData.AccountId,
			Type:      userData.Type,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, user)
	case "PATCH":
		fmt.Println("PATCH method not implemented")
	case "DELETE":
		fmt.Println("DELETE method not implemented")
	}
	return c.JSON(200, "ok")
}

func (uc UserController) InitRoutes() {
	uc.server.POST("/users", uc.ExecRoute)
	uc.server.GET("/users/:id", uc.ExecRoute)
	uc.server.PATCH("/users/:id", uc.ExecRoute)
	uc.server.DELETE("/users/:id", uc.ExecRoute)
}
