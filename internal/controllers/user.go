package controllers

import (
	"events-system/internal/services"
	usecases "events-system/internal/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	server  *echo.Echo
	useCase *usecases.UserUseCase
}

type UserDataDTO struct {
	Username string `json:"username" form:"username" query:"username"`
}

func NewUserController(server *echo.Echo, useCase *usecases.UserUseCase) *UserController {
	return &UserController{
		server:  server,
		useCase: useCase,
	}
}

func (uc UserController) ExecRoute(c echo.Context) error {
	fmt.Println(c.Request().Method, c.Path())
	switch method := c.Request().Method; method {
	case "GET":
		fmt.Println("GET METHOD")
		id := c.Param("id")
		fmt.Println(id)
		user, err := uc.useCase.GetUser(id)

		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		return c.JSON(200, user)
	case "POST":
		fmt.Println("start post method")
		userData := new(UserDataDTO)
		err := c.Bind(userData)
		if err != nil || len(userData.Username) == 0 {
			return c.String(http.StatusBadRequest, "bad request")
		}

		// TODO: fix, use only controller or usecase types
		user, err := uc.useCase.CreateUser(services.UserData{
			Username: userData.Username,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, user)
	case "PATCH":
		fmt.Println("PATCH method")
	case "DELETE":
		fmt.Println("DELETE method")
	}
	return c.JSON(200, "ok")
}

func (uc UserController) InitRoutes() {
	uc.server.POST("/users", uc.ExecRoute)
	uc.server.GET("/users/:id", uc.ExecRoute)
	uc.server.PATCH("/users/:id", uc.ExecRoute)
	uc.server.DELETE("/users/:id", uc.ExecRoute)
}
