package repository

type Task interface {
	FindByName(name string) entity.Task
}

import "punteros/domain/entity"
