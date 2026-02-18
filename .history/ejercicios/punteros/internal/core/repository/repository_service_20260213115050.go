package repository

import (

	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

type inMemoryTaskRepository struct {
	tasks map[uint8]domain.TaskType
}

// Add implements [ports.TaskRepository].
func (t *inMemoryTaskRepository) Add(task domain.TaskType) (domain.TaskType, error) {
	t.tasks[task.Id]=task
	return t.tasks[task.Id], nil
}

// Delete implements [ports.TaskRepository].
func (t *inMemoryTaskRepository) Delete(id uint8) {
	delete(t.tasks, id)
}

// FindAll implements [ports.TaskRepository].
func (t *inMemoryTaskRepository) FindAll() (map[uint8]domain.TaskType, error) {
	return t.tasks, nil
}

// FindById implements [ports.TaskRepository].
func (t *inMemoryTaskRepository) FindById(id uint8) (domain.TaskType, error) {
	if task, ok := t.tasks[id]; ok {
		return task, nil
	}
	return domain.TaskType{}, nil
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
