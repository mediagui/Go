package impl

import (
	t "punteros/domain/entity"
)

type taskType t.Task

var task taskType

type taskRepositoryImpl struct {
	task t.Task
}

func (*taskRepositoryImpl) addTask(task taskType)    {}
func (*taskRepositoryImpl) deleteTask(task taskType) {}
func (*taskRepositoryImpl) updateTask(task taskType) {}
func (*taskRepositoryImpl) findTaskById(id uint8)    {}

func NewTaskRepository(pTask t.Task) *taskRepositoryImpl {
	return &taskRepositoryImpl{
		task: pTask,
	}
}
