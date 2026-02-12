package repository

type Task struct {
	// Add Task fields here
}

type Task interface {
	FindByName(name string) Task
}
