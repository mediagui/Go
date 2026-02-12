package repository

import task "punteros/domain/entity"

type FindTask interface {
	ByName(name string) (task.Task, error)
}
type AddTask interface {
	New(task task.Task) (task.Task, error)
}
type UpdateTask interface {
	Update(task task.Task) (task.Task, error)
}

type DeleteTask interface {
	Delete(name string) bool
}
