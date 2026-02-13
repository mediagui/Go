// Struct to hold the task status, title and description
package domain

// Private struct to represent the task
type task struct {
	Id          uint8
	Title       string
	Description string
	status      TaskStatus
}

// Public type to define tasks
type TaskType task

func (s TaskStatus) String() string {
	return [...]string{"Completed", "Uncompleted"}[s]
}

func (s TaskStatus) IsCompleted() bool {
	return s == Completed
}
