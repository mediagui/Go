package service

import ports "punteros/internal/core/ports"

type taskService struct {
	repo ports.TaskRepository
}

func NewOrderService(repo ports.TaskRepository) ports.TaskService {

	return &taskService{
		repo: repo,
	}
}
