package application

import u "punteros/domain/usecase"

type taskApplicationImpl struct {
	taskUseCaseImpl u.TaskUseCase
}

type TaskApplication interface {
	TaskApplication() taskApplicationImpl
}

func NewApplication(pTaskUseCase u.TaskUseCase) *taskApplicationImpl {
	return &taskApplicationImpl{
		taskUseCaseImpl: pTaskUseCase,
	}
}
