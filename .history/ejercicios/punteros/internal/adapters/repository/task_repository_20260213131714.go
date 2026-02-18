// Package repository implementa los ADAPTADORES de persistencia.
// Los adaptadores son implementaciones concretas de los puertos (interfaces).
// Este paquete contiene toda la lógica específfica de cómo se almacenan los datos:
// - En este caso: en memoria (map)
// - En un proyecto real podría ser: SQL, MongoDB, archivos, etc
// VENTAJA: Si mañana cambias a una base de datos real, este es el único archivo
// que necesitas reescribir. El resto de la aplicación no cambia.
package repository

import (
	"punteros/internal/core/domain"
	"punteros/internal/core/ports"
)

// inMemoryTaskRepository es un ADAPTADOR que implementa la interfaz ports.TaskRepository.
// ENCAPSULACIÓN: El nombre comienza con minúscula para mantenerlo privado.
// Usuarios del paquete no conocen detalles sobre cómo se guarda (en memoria),
// solo interactuar a través de la interfaz TaskRepository.
// ALMACENAMIENTO: Usa un map[uint8]domain.TaskType para guardar tareas en memoria
type inMemoryTaskRepository struct {
	// tasks es el almacenamiento en memoria
	// Clave: ID de la tarea (uint8)
	// Valor: La estructura de la tarea (domain.TaskType)
	tasks map[uint8]domain.TaskType
}

// Add implementa el método Add de la interfaz TaskRepository.
// Añade una nueva tarea al mapa en memoria.
// Este método es responsable de:
// 1. Asignar un ID único a la tarea
// 2. Guardarla en el almacenamiento
// 3. Retornar la tarea guardada (con ID asignado)
func (t *inMemoryTaskRepository) Add(task domain.TaskType) (domain.TaskType, error) {
	// Generar un ID único basado en la cantidad de elementos
	task.Id = generateTaskId(uint8(len(t.tasks)))
	// Guardar en el mapa
	t.tasks[task.Id] = task
	// Retornar la tarea con el ID asignado
	return t.tasks[task.Id], nil
}

// Delete implementa el método Delete de la interfaz TaskRepository.
// Elimina una tarea por su ID del mapa en memoria.
func (t *inMemoryTaskRepository) Delete(id uint8) {
	delete(t.tasks, id)
}

// FindAll implementa el método FindAll de la interfaz TaskRepository.
// Retorna todas las tareas almacenadas.
func (t *inMemoryTaskRepository) FindAll() (map[uint8]domain.TaskType, error) {
	return t.tasks, nil
}

// FindById implementa el método FindById de la interfaz TaskRepository.
// Busca y retorna una tarea por su ID.
// Si no existe, retorna un TaskType vacío (valor por defecto).
func (t *inMemoryTaskRepository) FindById(id uint8) (domain.TaskType, error) {
	// Intentar obtener la tarea del mapa
	if task, ok := t.tasks[id]; ok {
		return task, nil
	}
	// Si no existe, retornar un TaskType vacío
	return domain.TaskType{}, nil
}

// Update implementa el método Update de la interfaz TaskRepository.
// Actualiza una tarea existente en el mapa.
func (t *inMemoryTaskRepository) Update(task domain.TaskType) (domain.TaskType, error) {
	// Sobrescribir la tarea en el mapa (si existe, se actualiza; si no, se añade)
	t.tasks[task.Id] = task
	return task, nil
}

// NewTaskRepository es el CONSTRUCTOR del adaptador de repositorio.
// Retorna una instancia de inMemoryTaskRepository encapsulada en la interfaz TaskRepository.
// PATRÓN: En Go es común que los constructores retornen la interfaz, no la implementación.
// Esto permite:
// 1. Cambiar la implementación futura sin afectar el código que lo usa
// 2. Retornar diferentes implementaciones según necesidad
// EJEMPLO:
//     repo := repository.NewTaskRepository()  // Retorna ports.TaskRepository
//     // En tests, podrías hacer:
//     repo := &MockRepository{}  // Otro adaptador que implementa la interfaz
func NewTaskRepository() ports.TaskRepository {
	return &inMemoryTaskRepository{
		// Inicializar el mapa vacío
		tasks: make(map[uint8]domain.TaskType),
	}
}

// generateTaskId es una función privada (comienza con minúscula) que genera IDs únicos.
// Estrategia simple: Usar la cantidad de elementos como ID.
// En una aplicación real, usarías:
// - Base de datos autoincrementales
// - UUIDs
// - Secuencias
func generateTaskId(mapLen uint8) uint8 {
	if mapLen == 0 {
		// Primer elemento, ID = 0
		return 0
	} else {
		// IDs posteriores = cantidad de elementos actual
		return mapLen
	}
}
