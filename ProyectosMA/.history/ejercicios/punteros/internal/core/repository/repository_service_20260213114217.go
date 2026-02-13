package repository

import (
	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

type taskRepository struct {
	tasks map[string]*domain.TaskType
}

// Add implements [ports.TaskRepository].
func (t *taskRepository) Add(task domain.TaskType) (domain.TaskType, error) {
	task
}

// Delete implements [ports.TaskRepository].
func (t *taskRepository) Delete(id uint8) bool {
	panic("unimplemented")
}

// FindAll implements [ports.TaskRepository].
func (t *taskRepository) FindAll() ([]domain.TaskType, error) {
	panic("unimplemented")
}

// FindById implements [ports.TaskRepository].
func (t *taskRepository) FindById(id uint8) (domain.TaskType, error) {
	panic("unimplemented")
}

// Update implements [ports.TaskRepository].
func (t *taskRepository) Update(task domain.TaskType) (domain.TaskType, error) {
	panic("unimplemented")
}

func NewTaskRepository() ports.TaskRepository {
	return &taskRepository{
		tasks: make(map[string]*domain.TaskType),
	}
}
