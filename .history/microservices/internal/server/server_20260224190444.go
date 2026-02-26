package server

import (
	"github.com/labstack/echo/v4"
	"github.com/mediagui/go-microservies/internal/database"
)

type Server interface {
	Start() error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(echo *echo.Echo, db database.DatabaseClient) *EchoServer {
	return &EchoServer{
		echo: echo,
		DB:   db,
	}
}
