package infrastructure

import d "punteros/domain/entity"

type dataBase struct {
	taskDatabase map[uint8]d.TaskType
}

func (db *dataBase) Add(task d.TaskType) {

	task.Id = uint8(len(db.taskDatabase)) + 1

	db.taskDatabase[task.Id] = task
}

func (db *dataBase) Update(task d.TaskType) {

	db.taskDatabase[task.Id] = task
}

func (db *dataBase) Find(id uint8) (d.TaskType, bool) {

	if task := db.taskDatabase[id]; task == nil {
		return
	}

}
