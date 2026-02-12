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

func (t *taskUseCaseImpl) addTask(task t.TaskType) {
	t.taskRepositoryImpl.AddNewTask(task)	
}
func (t *taskUseCaseImpl) deleteTask(task t.TaskType) {
	t.taskRepositoryImpl.DeleteTask(task.Id)
}
func (*taskUseCaseImpl) updateTask(task t.TaskType) {

}
func (*taskUseCaseImpl) findTaskById(id uint8) {

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
