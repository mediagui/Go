package impl

import (
	t "punteros/domain/entity"
)

type taskRepositoryImpl struct {
	task t.TaskType
}

func (*taskRepositoryImpl) addTask(task t.TaskType)    {}
func (*taskRepositoryImpl) deleteTask(task t.TaskType) {}
func (*taskRepositoryImpl) updateTask(task t.TaskType) {}
func (*taskRepositoryImpl) findTaskById(id uint8)      {}

func NewTaskRepository(pTask t.TaskType) *taskRepositoryImpl {
	return &taskRepositoryImpl{
		task: pTask,
	}
}
