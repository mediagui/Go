// Struct to hold the task status, title and description
package task

// Type to define the task status
type taskStatus uint8

// Pseudo enumerated types to represent the task status
const (
	TaskCompleteStatus taskStatus = iota
	TaskUncompleteStatus
)

// Struct to represent the task
type Task struct {
	Id          uint8
	Title       string
	Description string
	status      taskStatus
}
