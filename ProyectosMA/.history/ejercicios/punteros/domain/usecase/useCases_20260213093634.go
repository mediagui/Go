package usecase

import r "punteros/domain/repository"

type TaskUseCase interface {
	TaskRepository() r.TaskRepository
}
