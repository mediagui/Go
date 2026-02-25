package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mediagui/go-microservies/internal/database"
	"github.com/mediagui/go-microservies/internal/models"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context)
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(echo *echo.Echo, db database.DatabaseClient) *EchoServer {
	server := &EchoServer{
		echo: echo,
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %v", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() {

}

func (s *EchoServer) Readiness(ctx echo.Context) error {

	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}

	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	health := models.Health{
		Status: "alive",
	}
	ctx.JSON(http.StatusOK, health)
}
