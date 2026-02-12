package application

import u "punteros/domain/usecase"

type taskApplicationImpl struct {
	taskUseCaseImpl u.TaskUseCase
}

func NewApplication(pTaskUseCase u.TaskUseCase) *taskApplicationImpl {
	return &taskApplicationImpl{
		taskUseCaseImpl: pTaskUseCase,
	}
}
