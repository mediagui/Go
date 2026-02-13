package http

import (
	"punteros/internal/core/ports"
	"punteros/internal/core/service"
)

type taskService struct {
	service service.TaskService
}

func NewTaksHandler(service ports.TaskService) {

}
