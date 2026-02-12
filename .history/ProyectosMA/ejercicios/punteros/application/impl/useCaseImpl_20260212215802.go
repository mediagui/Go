package impl

import (
	r "punteros/domain/repository"
	u "punteros/domain/usecase"
)

// Concrete implementation of TaskUseCase
type taskUseCaseImpl struct {
	taskRepositoryImpl r.TaskRepository
}

func (t *taskUseCaseImpl) addTask(task taskType) {
	t.taskRepositoryImpl.AddNewTask()
}
func (*taskUseCaseImpl) deleteTask(task taskType) {}
func (*taskUseCaseImpl) updateTask(task taskType) {}
func (*taskUseCaseImpl) findTaskById(id uint8)    {}

// Constructor function - creates a new TaskUseCase instance
func NewTaskUseCase(pTaskRepository r.TaskRepository) u.TaskUseCase {
	return &taskUseCaseImpl{
		taskRepositoryImpl: pTaskRepository,
	}
}

// Implementation of TaskRepository() method from TaskUseCase interface
func (t *taskUseCaseImpl) TaskRepository() r.TaskRepository {
	return t.taskRepositoryImpl
}
