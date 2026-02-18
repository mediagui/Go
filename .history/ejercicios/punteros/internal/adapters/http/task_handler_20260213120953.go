package http

import (
	"punteros/internal/core/ports"
)

type TaskService struct {
	service ports.TaskService
}

func NewTaksHandler(service ports.TaskService) *TaskService {

	return &TaskService{
		service: service,
	}

}
