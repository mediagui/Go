package repository

import (
	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

type taskRepository struct {
	tasks map[string]*domain.TaskType
}

func NewTaskRepository() ports.TaskRepository {
	return &taskRepository{
		tasks: make(map[string]*domain.TaskType),
	}
}
