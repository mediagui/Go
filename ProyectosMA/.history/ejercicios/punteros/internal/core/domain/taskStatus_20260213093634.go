package util

// Type to define the task status
type TaskStatus uint8

// Pseudo enumerated types to represent the task status
const (
	Completed TaskStatus = iota
	Uncompleted
)

func (s TaskStatus) String() string {
	return [...]string{"Completed", "Uncompleted"}[s]
}

func (s TaskStatus) IsCompleted() bool {
	return s == Completed
}
