package impl

import t "punteros/domain/entity"

type taskRepositoryImpl struct {
	taskRepository r.TaskRepository
}

type TaskRepository interface {
	FindTaskById(id uint8) (t.Task, error)
	AddNewTask(task t.Task) (t.Task, error)
	UpdateTask(task t.Task) (t.Task, error)
	DeleteTask(name string) bool
	FindAll() []t.Task
}
