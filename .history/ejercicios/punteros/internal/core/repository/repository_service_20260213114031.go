package repository

import "punteros/internal/core/domain"
import "punteros/internal/core/ports"

type taskRepository struct {
	tasks map[string]*domain.TaskType
}

func NewTaskRepository() ports.TaskRepository {

}
