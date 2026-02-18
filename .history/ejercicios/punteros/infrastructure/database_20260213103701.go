package infrastructure

import d "punteros/domain/entity"

type DataBase struct {
	taskDatabase map[uint8]d.TaskType
}

func (db *DataBase) Add(task d.TaskType) (d.TaskType, bool) {

	task.Id = uint8(len(db.taskDatabase)) + 1

	db.taskDatabase[task.Id] = task

	// Se supone que no habrá errores
	return db.taskDatabase[task.Id], true

}

func (db *DataBase) Update(task d.TaskType) {

	db.taskDatabase[task.Id] = task
}

func (db *DataBase) Find(id uint8) (d.TaskType, bool) {

	if task, ok := db.taskDatabase[id]; ok {
		return task, true
	}
	return d.TaskType{}, false

}

func (db *DataBase) Delete(task d.TaskType) {
	delete(db.taskDatabase, task.Id)
}

func NewDatabase() *DataBase {

	return &DataBase{
		taskDatabase: map[uint8]d.TaskType{},
	}

}
