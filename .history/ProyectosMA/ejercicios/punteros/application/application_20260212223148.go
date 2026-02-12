package application

import u "punteros/domain/usecase"

type taskApplicationImpl struct {
	taskUseCaseImpl u.TaskUseCase
}

type TaskApplication interface {
	TaskApplication() r.TaskRepository
}

func NewApplication(pTaskUseCase u.TaskUseCase) taskApplicationImpl {
	return &taskApplicationImpl{
		taskUseCaseImpl: pTaskUseCase,
	}
}
