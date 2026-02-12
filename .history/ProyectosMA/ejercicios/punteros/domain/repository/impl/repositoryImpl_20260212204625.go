package impl

import r "punteros/domain/repository"
import t "punteros/domain/entity"

type repository t.Task

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
