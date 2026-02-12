package impl

import (
	t "punteros/domain/entity"
)

type taskRepositoryImpl struct {
	task t.Task
}

func NewTaskRepository(pTask t.Task) {
	taskRepositoryImpl.task = pTask
}
