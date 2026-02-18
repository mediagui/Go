package main

import (
	"fmt"
	"punteros/internal/adapters/http"
	"punteros/internal/adapters/repository"
	"punteros/internal/core/domain"
	"punteros/internal/core/service"
)

func main() {

	//Inicializamos e inyectamos las dependencias
	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository)
	taskHandler := http.NewTaskHandler(taskService)

	addTasks(taskHandler)

	updateTask(taskHandler)

	printTasks(taskHandler)

	deleteTask(taskHandler)

	printTasks(taskHandler)

}

func deleteTask(taskHandler *http.TaskHandler) {
	taskHandler.DeleteTask(1)
	println("Tarea borrada")
}

func updateTask(taskHandler *http.TaskHandler) {
	task, _ := taskHandler.FindTask(1)
	type status domain.TaskStatus
	task.Title = "Inicial - 1"
	task.Description = "Description - 1"
	task.Status = domain.Completed

	taskHandler.UpdateTask(task)

	println("Tarea 1 actualizada")
}

func addTasks(taskHandler *http.TaskHandler) {
	taskHandler.CreateTask(domain.TaskType{
		Title:       "Inicial",
		Description: "Descripión de la tarea",
		Status:      domain.Uncompleted,
	})

	println("Tarea 1 creada")

	taskHandler.CreateTask(domain.TaskType{
		Title:       "Inicial-2",
		Description: "Descripión de la tarea 2",
		Status:      domain.Uncompleted,
	})

	allTasks, _ := taskHandler.FindAllTasks()
	fmt.Println(allTasks)

	println("Tarea 2 creada")
}

func printTasks(taskHandler *http.TaskHandler) {
	allTasks, _ := taskHandler.FindAllTasks()
	fmt.Println(allTasks)
}
