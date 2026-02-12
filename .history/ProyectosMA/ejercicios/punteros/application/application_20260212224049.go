package application

import (
	u "punteros/domain/usecase"
	d "punteros/domain"
)

type taskApplicationImpl struct {
	taskUseCaseImpl u.TaskUseCase
}

func (t *taskApplicationImpl) AddNewTask(task d.TaskType) {
	t.TaskApplication().taskUseCaseImpl.TaskRepository().AddNewTask(task)
}
func (t *taskApplicationImpl) DeleteTask(task t.TaskType) {
	t.NewApplication().DeleteTask(task.Id)
}
func (t *taskApplicationImpl) UpdateTask(task t.TaskType) {
	t.NewApplication().UpdateTask(task)
}
func (t *taskApplicationImpl) FindTaskById(id uint8) {
	t.TaskApplication().FindTaskById(id)
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
