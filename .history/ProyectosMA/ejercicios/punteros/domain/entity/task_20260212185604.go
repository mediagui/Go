// Struct to hold the task status, title and description
package task

// Pseudo enumerated types to represent the task status
const (
	completeStatus uint8 = iota
	uncompleteStatus
)

// Type to define the task status
type taskStatus uint8

// Struct to represent the task
type Task struct {
	Id          uint8
	Title       string
	Description string
	status      taskStatus
}
