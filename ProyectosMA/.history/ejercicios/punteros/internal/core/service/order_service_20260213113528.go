package service

import (
	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

type taskService struct {
	repository ports.TaskRepository
}

func (service *taskService) AddTask(task domain.TaskType) (domain.TaskType, error) {
	return service.repository.Add(task)
}

func (service *taskService) DeleteTask(id uint8) bool {
	return service.repository.Delete(id)
}

func (service *taskService) FindAllTask() ([]domain.TaskType, error) {
	return service.repository.FindAll()
}

func (service *taskService) FindByIdTask(id uint8) (domain.TaskType, error) {
	return service.repository.FindById(id)
}

func (service *taskService) UpdateTask(task domain.TaskType) (domain.TaskType, error) {
	return service.repository.Update(task)
}

// Constructor
func NewOrderService(repo ports.TaskRepository) ports.TaskService {
	return &taskService{
		repository: repo,
	}
}
