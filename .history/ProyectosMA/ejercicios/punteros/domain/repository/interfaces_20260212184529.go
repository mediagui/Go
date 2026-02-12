package repository

import "punteros/domain/entity"

type FindTask interface {
	FindByName(name string) task.Task
}
