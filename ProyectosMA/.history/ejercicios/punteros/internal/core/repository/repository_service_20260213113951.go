package repository

import "punteros/internal/core/domain"

type taskRepository struct {
	tasks map[string]*domain.TaskType
}

func NewTaskRepository() {}
