package application

import (
	d "punteros/domain"
	u "punteros/domain/usecase"
)

type taskApplicationImpl struct {
	taskUseCaseImpl u.TaskUseCase
}

func (t *taskApplicationImpl) AddNewTask(task d.TaskType) {
	t.TaskApplication().taskUseCaseImpl.TaskRepository().AddNewTask(task)
}
func (t *taskApplicationImpl) DeleteTask(task d.TaskType) {
	t.TaskApplication().taskUseCaseImpl.TaskRepository().DeleteTask(task)
}
func (t *taskApplicationImpl) UpdateTask(task d.TaskType) {
	t.TaskApplication().taskUseCaseImpl.TaskRepository().UpdateTask(task)
}
func (t *taskApplicationImpl) FindTaskById(id uint8) {
	t.TaskApplication().taskUseCaseImpl.TaskRepository().FindTaskById(id)
}

type TaskApplication interface {
	TaskApplication() taskApplicationImpl
}

func NewTaskApplication(pTaskUseCase u.TaskUseCase) *taskApplicationImpl {
	return &taskApplicationImpl{
		taskUseCaseImpl: pTaskUseCase,
	}
}

func (t *taskApplicationImpl) TaskApplication() taskApplicationImpl{
	return t.TaskApplication()
}
