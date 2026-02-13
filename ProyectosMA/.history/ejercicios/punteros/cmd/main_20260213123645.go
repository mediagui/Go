package main

import (
	"fmt"
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
		Status:      domain.Uncompleted,
	})

	taskHandler.CreateTask(domain.TaskType{
		Title:       "Inicial-2",
		Description: "Descripión de la tarea 2",
		Status:      domain.Uncompleted,
	})

	allTasks, _ := taskHandler.FindAllTasks()
	fmt.Println(allTasks)

	task, _ := taskHandler.FindTask(1)

	type status domain.TaskStatus

	task.Title = "Inicial - 1"
	task.Description = "Description - 1"
	task.Status = domain.Completed

	taskHandler.UpdateTask(task)

	allTasks, _ = taskHandler.FindAllTasks()
	fmt.Println(allTasks)

}
