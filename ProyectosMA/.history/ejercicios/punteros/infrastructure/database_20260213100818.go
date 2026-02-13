package infrastructure

import d "punteros/domain/entity"

type dataBase struct {
	taskDatabase []d.TaskType
}

func (db *dataBase) Add(task d.TaskType) {
	db.taskDatabase = append(db.taskDatabase, task)
}
