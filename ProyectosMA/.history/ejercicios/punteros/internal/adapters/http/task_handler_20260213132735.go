// Package http implementa el ADAPTADOR de entrada HTTP (controlador).
// Este adaptador:
// - Recibe peticiones de la aplicación (en este caso, directas desde main)
// - Las traduce al formato que entiende el servicio
// - Llama al servicio con las dependencias inyectadas
// - Traduce las respuestas al formato esperado
// En una aplicación real con servidor HTTP, este archivo contendría
// los handlers para rutas como POST /tasks, GET /tasks/:id, etc.
// NOTA: Es un "adaptador" porque está adaptando el protocolo HTTP
// (o en este caso, llamadas directo) a las operaciones del dominio.
package http

import (
	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

// TaskHandler es un ADAPTADOR que traduce operaciones HTTP (o de entrada) a llamadas de servicio.
// INYECCIÓN DE DEPENDENCIAS: El servicio es inyectado en el constructor.
// TaskHandler depende de la INTERFAZ TaskService, no de la implementación concreta.
// Esto permite:
// - Cambiar la lógica de negocio sin cambiar el handler
// - Testear el handler con un mock del servicio
// - Añadir diferentes transportes (gRPC, WebSocket, etc) sin cambiar el servicio
type TaskHandler struct {
	// service es un campo privado que guarda la referencia al puerto TaskService
	// Es la INTERFAZ, permitiendo desacoplamiento total
	service ports.TaskService
}

// TaskHandlerPort es la interfaz que expone las operaciones del handler.
// Extraer esta interfaz permite a los callers depender de la abstracción
// en lugar de la implementación concreta `TaskHandler`.
type TaskHandlerPort interface {
	CreateTask(task domain.TaskType) (domain.TaskType, error)
	DeleteTask(id uint8)
	UpdateTask(task domain.TaskType) (domain.TaskType, error)
	FindTask(id uint8) (domain.TaskType, error)
	FindAllTasks() (map[uint8]domain.TaskType, error)
}

// CreateTask maneja la creación de nuevas tareas.
// Delega la lógica de negocio al servicio inyectado.
// En una aplicación real, este método sería llamado por un handler HTTP:
//
//	func (h *TaskHandler) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
//	    var task domain.TaskType
//	    // parsear JSON del request
//	    result, err := h.CreateTask(task)
//	    // retornar JSON response
//	}
func (h *TaskHandler) CreateTask(task domain.TaskType) (domain.TaskType, error) {
	// Delegar al servicio
	return h.service.AddTask(task)
}

// DeleteTask maneja la eliminación de tareas.
// Recibe el ID de la tarea a eliminar.
func (h *TaskHandler) DeleteTask(id uint8) {
	h.service.DeleteTask(id)
}

// UpdateTask maneja la actualización de tareas existentes.
// Recibe la tarea con los valores actualizados.
func (h *TaskHandler) UpdateTask(task domain.TaskType) (domain.TaskType, error) {
	return h.service.UpdateTask(task)
}

// FindTask busca una tarea específica por su ID.
func (h *TaskHandler) FindTask(id uint8) (domain.TaskType, error) {
	return h.service.FindByIdTask(id)
}

// FindAllTasks retorna todas las tareas disponibles.
func (h *TaskHandler) FindAllTasks() (map[uint8]domain.TaskType, error) {
	return h.service.FindAllTask()
}

// NewTaskHandler es el CONSTRUCTOR del adaptador HTTP.
// Realiza INYECCIÓN DE DEPENDENCIAS del servicio.
// PATRÓN EN GO: Los constructores reciben las dependencias como parámetros
// y retornan un puntero a la estructura.
// VENTAJAS:
// 1. Explícito: Se ve claramente qué necesita el handler
// 2. Testeable: Se puede pasar un mock del servicio
// 3. Flexible: Diferentes instancias pueden tener diferentes servicios
// EJEMPLO DE USO (en main.go):
//
//	service := service.NewTaskService(repo)
//	handler := http.NewTaskHandler(service)  // <- Inyectando el servicio
func NewTaskHandler(service ports.TaskService) TaskHandlerPort {
	return &TaskHandler{
		service: service,
	}
}
