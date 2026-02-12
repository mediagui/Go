package impl

import r "punteros/domain/repository"

type taskRepositoryImpl struct {
	taskRepository r.TaskRepository
}

func NewTaskRepository(taskRepository r.TaskRepository) r.TaskRepository {
	-
}

type TaskRepository interface {
	FindTaskById(id uint8) (t.Task, error)
	AddNewTask(task t.Task) (t.Task, error)
	UpdateTask(task t.Task) (t.Task, error)
	DeleteTask(name string) bool
	FindAll() []t.Task
}
