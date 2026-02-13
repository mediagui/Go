// Struct to hold the task status, title and description
package task

import u "punteros/domain/util"

// Struct to represent the task
type task struct {
	Id          uint8
	Title       string
	Description string
	status      u.TaskStatus
}

type TaskType task

// var task taskType
