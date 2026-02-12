package application

import u "punteros/domain/usecase"

type taskApplicationImpl struct {
	taskUseCaseImpl u.TaskUseCase
}

func (t *taskApplicationImpl) AddNewTask(task t.TaskType) {
	t.NewApplication().AddNewTask(task)
}
func (t *taskApplicationImpl) DeleteTask(task t.TaskType) {
	t.NewApplication().DeleteTask(task.Id)
}
func (t *taskApplicationImpl) UpdateTask(task t.TaskType) {
	t.NewApplication().UpdateTask(task)
}
func (t *taskApplicationImpl) FindTaskById(id uint8) {
	t.NewApplication().FindTaskById(id)
}

type TaskApplication interface {
	TaskApplication() taskApplicationImpl
}

func NewApplication(pTaskUseCase u.TaskUseCase) *taskApplicationImpl {
	return &taskApplicationImpl{
		taskUseCaseImpl: pTaskUseCase,
	}
}
