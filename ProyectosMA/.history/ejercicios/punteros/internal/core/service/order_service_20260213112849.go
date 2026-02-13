package service

import ports "punteros/internal/core/ports"

type taskService struct {
	repo ports.TaskRepository
}

func NewOrderService(repository ports.TaskRepository) ports.TaskService {

	return &taskService{
		repo: repository,
	}
}
