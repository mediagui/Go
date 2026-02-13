package ports

import "punteros/internal/core/domain"

type TaskRepository interface {
	FindById(id uint8) (domain.TaskType, error)
	Add(task domain.TaskType) (domain.TaskType, error)
	Update(task domain.TaskType) (domain.TaskType, error)
	Delete(id uint8)
	FindAll() (map[uint8]domain.TaskType, error)
}

type TaskService interface {
	FindByIdTask(id uint8) (domain.TaskType, error)
	AddTask(task domain.TaskType) (domain.TaskType, error)
	UpdateTask(task domain.TaskType) (domain.TaskType, error)
	DeleteTask(id uint8)
	FindAllTask() (map[uint8]domain.TaskType, error)
}
