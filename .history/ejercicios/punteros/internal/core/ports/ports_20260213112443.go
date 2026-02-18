package ports

import t "punteros/internal/core/domain"

type TaskRepository interface {
	FindById(id uint8) (t.TaskType, error)
	Add(task t.TaskType) (t.TaskType, error)
	Update(task t.TaskType) (t.TaskType, error)
	Delete(id uint8) bool
	FindAll() ([]t.TaskType, error)
}

type TaskService interface {
	FindByIdTask(id uint8) (t.TaskType, error)
	AddTask(task t.TaskType) (t.TaskType, error)
	UpdateTask(task t.TaskType) (t.TaskType, error)
	DeleteTask(id uint8) bool
	FindAllTask() ([]t.TaskType, error)
}
