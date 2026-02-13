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

func (h *TaskHandler) DeleteTask(id uint8) {
	h.service.DeleteTask(id)
}

func (h *TaskHandler) UpdateTask(task domain.TaskType) (domain.TaskType, error) {
	return h.service.UpdateTask(task)
}

func (h *TaskHandler) FindTask(id uint8) (domain.TaskType, error) {
	return h.service.FindByIdTask(id)
}

func (h *TaskHandler) FindAllTasks() (map[uint8]domain.TaskType, error) {
	return h.service.FindAllTask()
}

func NewTaksHandler(service ports.TaskService) *TaskHandler {

	return &TaskHandler{
		service: service,
	}

}
