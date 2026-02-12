package impl

import (
	t "punteros/domain/entity"
)

var task *t.Task = &t.Task{}

type taskRepositoryImpl struct {
	task t.Task
}

func (*t.Task) addTask() {}

func NewTaskRepository(pTask t.Task) *taskRepositoryImpl {
	return &taskRepositoryImpl{
		task: pTask,
	}
}
