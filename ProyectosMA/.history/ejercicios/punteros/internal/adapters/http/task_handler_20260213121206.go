package http

import (
	"punteros/internal/core/ports"
)

type TaskHandler struct {
	service ports.TaskService
}

func (h *TaskHandler) CreateTask()

func NewTaksHandler(service ports.TaskService) *TaskService {

	return &TaskHandler{
		service: service,
	}

}
