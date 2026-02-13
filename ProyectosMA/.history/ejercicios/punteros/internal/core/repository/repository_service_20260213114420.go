package repository

import (
	"strconv"

	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

type inMemoryTaskRepository struct {
	tasks map[string]*domain.TaskType
}

// Add implements [ports.TaskRepository].
func (t *inMemoryTaskRepository) Add(task domain.TaskType) (domain.TaskType, error) {
	t.tasks[strconv.FormatUint(uint64(task.Id), 10)] = &task
	return task, nil
}

// Delete implements [ports.TaskRepository].
func (t *inMemoryTaskRepository) Delete(id uint8) bool {
	panic("unimplemented")
}

// FindAll implements [ports.TaskRepository].
func (t *inMemoryTaskRepository) FindAll() ([]domain.TaskType, error) {
	panic("unimplemented")
}

// FindById implements [ports.TaskRepository].
func (t *inMemoryTaskRepository) FindById(id uint8) (domain.TaskType, error) {
	panic("unimplemented")
}

// Update implements [ports.TaskRepository].
func (t *inMemoryTaskRepository) Update(task domain.TaskType) (domain.TaskType, error) {
	panic("unimplemented")
}

func NewTaskRepository() ports.TaskRepository {
	return &inMemoryTaskRepository{
		tasks: make(map[string]*domain.TaskType),
	}
}
