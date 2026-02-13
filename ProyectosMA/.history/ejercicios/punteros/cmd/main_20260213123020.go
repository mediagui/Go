package main

import (
	"punteros/internal/adapters/http"
	"punteros/internal/adapters/repository"
	"punteros/internal/core/domain"
	"punteros/internal/core/service"
)

func main() {

	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository)
	taskHandler := http.NewTaskHandler(taskService)

	taskHandler.CreateTask(domain.TaskType{
		Title:       "Inicial",
		Description: "Descripión de la tarea",
	})

}
