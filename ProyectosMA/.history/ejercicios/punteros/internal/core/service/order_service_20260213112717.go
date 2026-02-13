package service

import ports "punteros/internal/core/ports"

type taskService struct {
	repo ports.TaskRepository
}
