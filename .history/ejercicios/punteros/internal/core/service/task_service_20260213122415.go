package service

import (
	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

type TaskService struct {
	repository ports.TaskRepository
}

func (service *TaskService) AddTask(task domain.TaskType) (domain.TaskType, error) {
	return service.repository.Add(task)
}

func (service *TaskService) DeleteTask(id uint8) {
	service.repository.Delete(id)
}

func (service *TaskService) FindAllTask() (map[uint8]domain.TaskType, error) {
	return service.repository.FindAll()
}

func (service *TaskService) FindByIdTask(id uint8) (domain.TaskType, error) {
	return service.repository.FindById(id)
}

func (service *TaskService) UpdateTask(task domain.TaskType) (domain.TaskType, error) {
	return service.repository.Update(task)
}

// Constructor
func NewTaskService(repo ports.TaskRepository) ports.TaskService {
	return &TaskService{
		repository: repo,
	}
}
