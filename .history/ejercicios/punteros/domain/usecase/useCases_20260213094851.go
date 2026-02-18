package usecase

import r "punteros/domain/repository"
import d "punteros/domain/entity"

// type TaskUseCase interface {
// 	TaskRepository() r.TaskRepository
// }

type taskUseCaseImpl struct {
	taskRepositoryImpl r.TaskRepository
}

func (t *taskUseCaseImpl) AddNewTask(task d.TaskType) {
	t.TaskApplication().taskUseCaseImpl.TaskRepository().AddNewTask(task)
}
func (t *taskUseCaseImpl) DeleteTask(task d.TaskType) {
	t.TaskApplication().taskUseCaseImpl.TaskRepository().DeleteTask(task.Id)
}
func (t *taskUseCaseImpl) UpdateTask(task d.TaskType) {
	t.TaskApplication().taskUseCaseImpl.TaskRepository().UpdateTask(task)
}
func (t *taskUseCaseImpl) FindTaskById(id uint8) {
	t.TaskApplication().taskUseCaseImpl.TaskRepository().FindTaskById(id)
}


func (t *Tas)
