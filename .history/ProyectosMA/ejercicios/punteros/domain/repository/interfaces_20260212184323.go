package repository

type Task struct {
	// Add Task fields here
}

type FindTask interface{
	FindByName(name string) Task

}

