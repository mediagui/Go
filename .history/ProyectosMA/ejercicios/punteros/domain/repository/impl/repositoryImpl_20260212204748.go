package impl

import (
	t "punteros/domain/entity"
	r "punteros/domain/repository"

	u "punteros/domain/util"
)

// Struct to represent the task
type Task struct {
	Id          uint8
	Title       string
	Description string
	status      u.TaskStatus
}

func NewTaskRepository() r.TaskRepository {
	return nil
}

type TaskRepository interface {
	FindTaskById(id uint8) (t.Task, error)
	AddNewTask(task t.Task) (t.Task, error)
	UpdateTask(task t.Task) (t.Task, error)
	DeleteTask(name string) bool
	FindAll() []t.Task
}
