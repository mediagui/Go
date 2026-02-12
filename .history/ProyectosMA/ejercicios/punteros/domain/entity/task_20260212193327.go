// Struct to hold the task status, title and description
package task

// Struct to represent the task
type Task struct {
	Id          uint8
	Title       string
	Description string
	status      taskStatus
}
