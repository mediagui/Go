package usecase

import r "punteros/domain/repository"

// type TaskRepository interface {
// 	FindTaskById(id uint8) (t.Task, error)
// 	AddNewTask(task t.Task) (t.Task, error)
// 	UpdateTask(task t.Task) (t.Task, error)
// 	DeleteTask(name string) bool
// 	FindAll() []t.Task
// }


type FindTaskUseCase interface{
	TaskRepository r.TaskRepository
}
