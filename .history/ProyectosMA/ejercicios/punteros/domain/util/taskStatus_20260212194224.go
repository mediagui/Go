package util

// Type to define the task status
type TaskStatus uint8

// Pseudo enumerated types to represent the task status
const (
	Completed TaskStatus = iota
	Uncompleted
)
