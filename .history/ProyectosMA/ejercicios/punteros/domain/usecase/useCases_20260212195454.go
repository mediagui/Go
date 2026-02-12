package usecase

import "punteros/domain"

type TaskRepository interface {
	FindTaskById(id uint8) (domain.Task, error)
	AddNewTask(task domain.Task) (domain.Task, error)
	UpdateTask(task domain.Task) (domain.Task, error)
	DeleteTask(name string) bool
	FindAll() []domain.Task
}

type FindTaskUseCase interface{

}
