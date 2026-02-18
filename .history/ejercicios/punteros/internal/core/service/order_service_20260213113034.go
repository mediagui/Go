package service

import (
	"punteros/internal/core/domain"
	ports "punteros/internal/core/ports"
)

type taskService struct {
	repo ports.TaskRepository
}

// AddTask implements [ports.TaskService].
func (t *taskService) AddTask(task domain.TaskType) (domain.TaskType, error) {
	return ports.TaskRepository.Add(task)
}

// DeleteTask implements [ports.TaskService].
func (t *taskService) DeleteTask(id uint8) bool {
	panic("unimplemented")
}

// FindAllTask implements [ports.TaskService].
func (t *taskService) FindAllTask() ([]domain.TaskType, error) {
	panic("unimplemented")
}

// FindByIdTask implements [ports.TaskService].
func (t *taskService) FindByIdTask(id uint8) (domain.TaskType, error) {
	panic("unimplemented")
}

// UpdateTask implements [ports.TaskService].
func (t *taskService) UpdateTask(task domain.TaskType) (domain.TaskType, error) {
	panic("unimplemented")
}

func NewOrderService(repo ports.TaskRepository) ports.TaskService {

	return &taskService{
		repo: repo,
	}
}
