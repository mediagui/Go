package repository

import t "punteros/domain/entity"

type TaskRepository interface {
	FindTaskByName(name string) (t.Task, error)
	AddNewTask(task t.Task) (t.Task, error)
	UpdateTask(task t.Task) (t.Task, error)
	DeleteTask(name string) bool
}
