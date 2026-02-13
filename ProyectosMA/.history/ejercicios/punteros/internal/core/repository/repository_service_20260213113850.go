package repository

type taskRepository struct {
	orders map[string]*domain.Order
}
