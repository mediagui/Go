package application

import r "punteros/domain/repository"

// Concrete implementation of TaskUseCase
type taskUseCaseImpl struct {
	taskRepository r.TaskRepository
}

// Constructor function - creates a new TaskUseCase instance
func NewTaskUseCase(taskRepository r.TaskRepository) TaskUseCase {
	return &taskUseCaseImpl{
		taskRepository: taskRepository,
	}
}

// Implementation of TaskRepository() method from TaskUseCase interface
func (t *taskUseCaseImpl) TaskRepository() r.TaskRepository {
	return t.taskRepository
}
