package dberrors

type NotFoundError struct {
	Entity string
	ID     string
}

func (e NotFoundError) Error() string {
	return "not found error: entity " + e.Entity + " with ID " + e.ID
}
