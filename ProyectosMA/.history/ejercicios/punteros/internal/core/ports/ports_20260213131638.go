// Package ports define las INTERFACES (puertos) que son el corazón del patrón hexagonal.
// Las interfaces actúan como contratos que separan:
// - El dominio/negocio (qué se debe hacer) de
// - Los adaptadores (cómo se hace)
//
// VENTAJAS DEL USO DE INTERFACES COMO PUERTOS:
// 1. Desacoplamiento: El servicio no conoce detalles específicos del repositorio
// 2. Testabilidad: Se puede crear un mock del repositorio fácilmente
// 3. Flexibilidad: Cambiar de base de datos sin modificar la lógica de negocio
// 4. Inyección de Dependencias: Las interfaces específican qué se necesita inyectar
package ports

import "punteros/internal/core/domain"

// TaskRepository es un PUERTO (interfaz) que define el contrato para la persistencia de tareas.
// Representa todas las operaciones de lectura/escritura que la aplicación necesita realizar.
// Los adaptadores concretos (en memoria, SQL, MongoDB, etc) implementan esta interfaz.
// Si mañana quieres cambiar de base de datos, sólo necesitas:
// 1. Crear una nueva struct que implemente TaskRepository
// 2. Inyectarla en el servicio
// 3. El resto de código no cambia
type TaskRepository interface {
	// FindById recupera una tarea por su ID
	// Retorna la tarea y un error si algo falla
	FindById(id uint8) (domain.TaskType, error)

	// Add añade una nueva tarea al repositorio
	// La implementación es responsable de asignar un ID
	// Retorna la tarea guardada (con ID asignado) y un error si falla
	Add(task domain.TaskType) (domain.TaskType, error)

	// Update modifica una tarea existente
	// Retorna la tarea actualizada y un error si falla
	Update(task domain.TaskType) (domain.TaskType, error)

	// Delete elimina una tarea por su ID
	Delete(id uint8)

	// FindAll retorna todas las tareas del repositorio
	// Se retorna un mapa para acceso rápido por ID
	FindAll() (map[uint8]domain.TaskType, error)
}

// TaskService es un PUERTO (interfaz) que define el contrato para la lógica de negocio.
// Representa los casos de uso principales de la aplicación.
// La implementación concreta (TaskService struct) utiliza el repositorio para ejecutar operaciones.
// Ventaja: El handler HTTP depende de esta interfaz, no de la implementación,
// permitiendo cambiar la lógica de negocio o hacer mocks para tests.
type TaskService interface {
	// FindByIdTask busca una tarea por su ID
	FindByIdTask(id uint8) (domain.TaskType, error)

	// AddTask crea una nueva tarea
	AddTask(task domain.TaskType) (domain.TaskType, error)

	// UpdateTask modifica una tarea existente
	UpdateTask(task domain.TaskType) (domain.TaskType, error)

	// DeleteTask elimina una tarea
	DeleteTask(id uint8)

	// FindAllTask retorna todas las tareas
	FindAllTask() (map[uint8]domain.TaskType, error)
}
