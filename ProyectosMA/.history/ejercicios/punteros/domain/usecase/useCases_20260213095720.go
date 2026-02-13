package usecase

import d "punteros/domain/entity"

type TaskUseCase interface {
	AddNewTask(task d.TaskType)
	DeleteTask(task d.TaskType)
	UpdateTask(task d.TaskType)
	FindTaskById(id uint8)
}
