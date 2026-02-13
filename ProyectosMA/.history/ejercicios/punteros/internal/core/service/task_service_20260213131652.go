// Package service implementa la lógica de negocio de la aplicación.
// El servicio es la capa intermedia entre los adaptadores (handlers, controladores) y
// los adaptadores de persistencia (repositorio).
// RESPONSABILIDADES DEL SERVICIO:
// 1. Implementar los casos de uso (lógica de negocio)
// 2. Orquestar operaciones del repositorio
// 3. Aplicar reglas de negocio
// 4. Manejar excepciones y errores
pkg package service

import (
	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

// TaskService es la estructura que implementa la interfaz ports.TaskService.
// INYECCIÓN DE DEPENDENCIAS: El repository es inyectado mediante el constructor,
// no creado internamente. Esto permite:
// - Cambiar la implementación del repositorio sin modificar TaskService
// - Usar un mock del repositorio en tests unitarios
// - Desacoplar completamente la lógica de negocio de la persistencia
type TaskService struct {
	// repository es un campo privado que guarda una referencia a la interfaz TaskRepository
	// Note: Es la INTERFAZ, no la implementación concreta
	// Esto es lo que permite el desacoplamiento
	repository ports.TaskRepository
}

// AddTask implementa el método de la interfaz TaskService.
// Añade una nueva tarea delegando al repositorio.
// FLUJO: Handler → Service.AddTask → Repository.Add → Almacenamiento
func (service *TaskService) AddTask(task domain.TaskType) (domain.TaskType, error) {
	// Delega la operación al repositorio
	// El servicio no sabe cómo se almacena, solo que debe hacerse
	return service.repository.Add(task)
}

// DeleteTask implementa el método de la interfaz TaskService.
// Elimina una tarea delegando al repositorio.
func (service *TaskService) DeleteTask(id uint8) {
	service.repository.Delete(id)
}

// FindAllTask implementa el método de la interfaz TaskService.
// Retorna todas las tareas del repositorio.
func (service *TaskService) FindAllTask() (map[uint8]domain.TaskType, error) {
	return service.repository.FindAll()
}

// FindByIdTask implementan implementing el método de la interfaz TaskService.
// Busca una tarea por su ID.
func (service *TaskService) FindByIdTask(id uint8) (domain.TaskType, error) {
	return service.repository.FindById(id)
}

// UpdateTask implementa el método de la interfaz TaskService.
// Actualiza una tarea delegando al repositorio.
func (service *TaskService) UpdateTask(task domain.TaskType) (domain.TaskType, error) {
	return service.repository.Update(task)
}

// NewTaskService es el CONSTRUCTOR que realizar INYECCIÓN DE DEPENDENCIAS.
// Patrón básico en Go: constructores que retornan la interfaz, no la implementación.
// EJEMPLO DE USO (en main.go):
//     repo := repository.NewTaskRepository()
//     service := service.NewTaskService(repo)  // <- Inyectando el repositorio
// VENTAJAS:
// 1. Fácil de testear: Pasar un mock del repositorio
// 2. Fácil de cambiar: Cambiar la implementación del repositorio en un solo lugar
// 3. Explícito: El código muestra claramente las dependencias
func NewTaskService(repo ports.TaskRepository) ports.TaskService {
	return &TaskService{
		repository: repo,
	}
}
