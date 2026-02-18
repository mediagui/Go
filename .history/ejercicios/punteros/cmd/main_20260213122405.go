package main

import (
	"punteros/internal/adapters/repository"
	"punteros/internal/core/service"
)

func main() {

	taskRepo := repository.NewTaskRepository()
	taskService := service.NewOrderService()
}
