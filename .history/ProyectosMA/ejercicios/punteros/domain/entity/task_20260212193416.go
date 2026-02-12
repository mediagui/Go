// Struct to hold the task status, title and description
package task

import util

// Struct to represent the task
type Task struct {
	Id          uint8
	Title       string
	Description string
	status      TaskStatus
}
