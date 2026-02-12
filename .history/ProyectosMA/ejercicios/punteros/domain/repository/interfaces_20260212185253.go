package repository

import t "punteros/domain/entity"

type FindTask interface {
	ByName(name string) (t.Task, error)
}
type AddTask interface {
	New(task t.Task) (t.Task, error)
}

type UpdateTask interface {
	Update(task t.Task) (t.Task, error)
}

type DeleteTask interface {
	Delete(name string) bool
}
