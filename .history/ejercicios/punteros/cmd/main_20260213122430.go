package main

import (
	"punteros/internal/adapters/repository"
	"punteros/internal/core/service"
)

func main() {

	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository)

}
