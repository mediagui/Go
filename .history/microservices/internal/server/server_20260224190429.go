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
