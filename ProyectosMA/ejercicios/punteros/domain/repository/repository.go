package repository

import t "punteros/domain/entity"

type TaskRepository interface {
	FindTaskById(id uint8) (t.TaskType, error)
	AddNewTask(task t.TaskType) (t.TaskType, error)
	UpdateTask(task t.TaskType) (t.TaskType, error)
	DeleteTask(id uint8) bool
	FindAll() []t.TaskType
}
