package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	Instance *echo.Echo
}

func NewEchoInstance() *Server {
	return &Server{
		Instance: echo.New(),
	}
}

func (s Server) Start(port string) {
	s.Instance.Logger.Fatal(
		s.Instance.Start(":1488"),
	)
}
