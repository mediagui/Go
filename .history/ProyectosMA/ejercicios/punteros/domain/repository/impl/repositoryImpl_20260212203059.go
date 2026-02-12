package impl

import r "punteros/domain/repository"
import r "punteros/domain/repository"

type taskRepository struct{
	taskRepository r.TaskRepository
}


func NewTaskRepository[t.Task]() r.TaskRepository {

}




type TaskRepository interface {
	FindTaskById(id uint8) (t.Task, error)
	AddNewTask(task t.Task) (t.Task, error)
	UpdateTask(task t.Task) (t.Task, error)
	DeleteTask(name string) bool
	FindAll() []t.Task
}
