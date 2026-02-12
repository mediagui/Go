package impl

import (
	r "punteros/domain/repository"
	u "punteros/domain/usecase"
)

// Concrete implementation of TaskUseCase
type taskUseCaseImpl struct {
	taskRepository r.TaskRepository
}

// Constructor function - creates a new TaskUseCase instance
func NewTaskUseCase(pTaskRepository r.TaskRepository) u.TaskUseCase {
	return &taskUseCaseImpl{
		taskRepository: pTaskRepository,
	}
}

// Implementation of TaskRepository() method from TaskUseCase interface
func (t *taskUseCaseImpl) TaskRepository() r.TaskRepository {
	return t.taskRepository
}
