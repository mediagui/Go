// Struct to hold the task status, title and description
package task

import d "punteros/internal/core/domain"

// Private struct to represent the task
type task struct {
	Id          uint8
	Title       string
	Description string
	status      d.TaskStatus
}

// Public type to define tasks
type TaskType task
