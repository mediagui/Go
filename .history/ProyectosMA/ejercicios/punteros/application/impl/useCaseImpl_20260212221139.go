package impl

import (
	t "punteros/domain/entity"
	r "punteros/domain/repository"
	u "punteros/domain/usecase"
)

// Concrete implementation of TaskUseCase
type taskUseCaseImpl struct {
	taskRepositoryImpl r.TaskRepository
}

func (t *taskUseCaseImpl) AddTask(task t.TaskType) {
	t.taskRepositoryImpl.AddNewTask(task)
}
func (t *taskUseCaseImpl) DeleteTask(task t.TaskType) {
	t.taskRepositoryImpl.DeleteTask(task.Id)
}
func (t *taskUseCaseImpl) UpdateTask(task t.TaskType) {
	t.taskRepositoryImpl.UpdateTask(task)
}
func (t *taskUseCaseImpl) FindTaskById(id uint8) {
	t.taskRepositoryImpl.FindTaskById(id)
}

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
