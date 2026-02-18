package impl

import (
	t "punteros/domain/entity"
	db "punteros/infrastructure"
)

type taskRepositoryImpl struct {
	task     t.TaskType
	database db.DataBase
}

func (*taskRepositoryImpl) AddTask(task t.TaskType) {

}
func (*taskRepositoryImpl) DeleteTask(task t.TaskType) {}
func (*taskRepositoryImpl) UpdateTask(task t.TaskType) {}
func (*taskRepositoryImpl) FindTaskById(id uint8)      {}

func NewTaskRepository(pTask t.TaskType) *taskRepositoryImpl {
	return &taskRepositoryImpl{
		task: pTask,
	}
}
