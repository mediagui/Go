// Struct to hold the task status, title and description
package task

import u "punteros/domain/util"

// Private struct to represent the task
type task struct {
	Id          uint8
	Title       string
	Description string
	status      u.TaskStatus
}

// Public type to define tasks
type TaskType task
