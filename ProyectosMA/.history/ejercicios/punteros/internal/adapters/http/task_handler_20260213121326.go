package http

import (
	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

type TaskHandler struct {
	service ports.TaskService
}

func (h *TaskHandler) CreateTask(task domain.TaskType) (domain.TaskType, error) {
	return h.service.AddTask(task)
}

func NewTaksHandler(service ports.TaskService) *TaskService {

	return &TaskHandler{
		service: service,
	}

}
