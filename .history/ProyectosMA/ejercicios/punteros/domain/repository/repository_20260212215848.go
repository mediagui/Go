package repository

import t "punteros/domain/entity"

type TaskRepository interface {
	FindTaskById(id uint8) (t.Task, error)
	AddNewTask(task t.TaskType) (t.Task, error)
	UpdateTask(task t.Task) (t.Task, error)
	DeleteTask(name string) bool
	FindAll() []t.Task
}
