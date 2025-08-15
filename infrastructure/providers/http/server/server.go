package server

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Instance *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewEchoInstance() *Server {
	return &Server{
		Instance: echo.New(),
	}
}

func (s Server) Start(port string) {
	s.Instance.Validator = &CustomValidator{validator: validator.New()}
	s.Instance.Logger.Fatal(
		s.Instance.Start(":1488"),
	)
}

func (s Server) Close(ctx context.Context) error {
	if err := s.Instance.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}
