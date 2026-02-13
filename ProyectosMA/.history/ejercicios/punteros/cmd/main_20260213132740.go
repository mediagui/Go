// Package main - Punto de entrada de la aplicación
// Este archivo orquesta la inyección de dependencias del patrón hexagonal.
// Es responsable de:
// 1. Crear instancias de los adaptadores (repositorio, handler)
// 2. Inyectar las dependencias en el orden correcto
// 3. Orquestar los casos de uso de la aplicación
package main

import (
	"fmt"
	"punteros/internal/adapters/http"
	"punteros/internal/adapters/repository"
	"punteros/internal/core/domain"
	"punteros/internal/core/service"
)

// main es el punto de entrada de la aplicación
// Implementa el patrón de inyección de dependencias de abajo hacia arriba:
// 1. Crea el repositorio (adaptador de persistencia)
// 2. Crea el servicio, inyectando el repositorio
// 3. Crea el handler, inyectando el servicio
// Esta estrategia permite cambiar implementaciones sin afectar componentes superiores
func main() {

	// INYECCIÓN DE DEPENDENCIAS - Construcción de la cadena de dependencias
	// Orden importante: Adaptadores abajo (no dependen de nada) → Servicios → Handlers

	// 1. Crear el adaptador de persistencia (repositorio en memoria)
	taskRepository := repository.NewTaskRepository()

	// 2. Crear el servicio, inyectando el repositorio como dependencia
	// El servicio recibe la interfaz TaskRepository, no la implementación concreta
	// Esto permite cambiar la persistencia sin modificar el servicio
	taskService := service.NewTaskService(taskRepository)

	// 3. Crear el adaptador HTTP (handler), inyectando el servicio
	// El handler recibe la interfaz TaskService, permitiendo tests con mocks
	taskHandler := http.NewTaskHandler(taskService)

	// Ejecutar casos de uso de prueba
	addTasks(taskHandler)
	updateTask(taskHandler)
	printTasks(taskHandler)
	deleteTask(taskHandler)
	printTasks(taskHandler)

}

// deleteTask - Demuestra la operación DELETE en el CRUD
// Utiliza el handler inyectado para eliminar una tarea por su ID
func deleteTask(taskHandler http.TaskHandlerPort) {
	taskHandler.DeleteTask(1)
	println("Tarea borrada")
}

// updateTask - Demuestra la operación UPDATE en el CRUD
// Obtiene una tarea existente, modifica sus campos y la actualiza
func updateTask(taskHandler http.TaskHandlerPort) {
	// Obtener la tarea con ID 1
	task, _ := taskHandler.FindTask(1)
	type status domain.TaskStatus
	// Modificar los campos de la tarea
	task.Title = "Inicial - 1"
	task.Description = "Description - 1"
	task.Status = domain.Completed

	// Persisti través del handler → servicio → repositorio
	taskHandler.UpdateTask(task)

	println("Tarea 1 actualizada")
}

// addTasks - Demuestra la operación CREATE en el CRUD
// Crea nuevas tareas usando el handler inyectado
func addTasks(taskHandler http.TaskHandlerPort) {
	// Crear la primera tarea
	// El flujo es: handler → servicio → repositorio → almacenamiento
	taskHandler.CreateTask(domain.TaskType{
		Title:       "Inicial",
		Description: "Descripión de la tarea",
		Status:      domain.Uncompleted,
	})

	println("Tarea 1 creada")

	// Crear la segunda tarea
	taskHandler.CreateTask(domain.TaskType{
		Title:       "Inicial-2",
		Description: "Descripión de la tarea 2",
		Status:      domain.Uncompleted,
	})

	// Listar todas las tareas creadas
	allTasks, _ := taskHandler.FindAllTasks()
	fmt.Println(allTasks)

	println("Tarea 2 creada")
}

// printTasks - Demuestra la operación READ en el CRUD
// Lista todas las tareas desde el repositorio a través del handler
func printTasks(taskHandler http.TaskHandlerPort) {
	// Obtener y mostrar todas las tareas
	allTasks, _ := taskHandler.FindAllTasks()
	fmt.Println(allTasks)
}
