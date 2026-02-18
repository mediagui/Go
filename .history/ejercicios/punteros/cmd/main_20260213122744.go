package main

import (
	"punteros/internal/adapters/http"
	"punteros/internal/adapters/repository"
	"punteros/internal/core/service"
)

func main() {

	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository)
	_ = http.NewTaskHandler(taskService)

}
