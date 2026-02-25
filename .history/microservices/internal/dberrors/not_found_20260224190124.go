package dberrors

import "fmt"

type NotFoundError struct {
	Entity string
	ID     string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("not found error: entity %s with ID %s", e.Entity, e.ID)
}
