// Package domain contiene las entidades y tipos del dominio de negocio.
package domain

// TaskStatus es un tipo de enumeración que representa los posibles estados de una tarea.
// En Go, usamos un tipo base (uint8) y constantes para crear enumeraciones pseudo-tipadas.
// Esto proporciona:
// - Type safety: No se pueden pasar integers arbitrarios donde se espera TaskStatus
// - Claridad: El código es explícito sobre qué representa cada valor
// - Métodos personalizados: Podemos definir métodos específicos para TaskStatus
type TaskStatus uint8

// Constantes enumeradas usando iota para generar valores secuenciales automáticamente.
// iota comienza en 0 y se incrementa para cada constante.
// Completed = 0 (completa/finalizada)
// Uncompleted = 1 (pendiente/incompleta)
const (
	// Completed representa una tarea que ha sido finalizada
	Completed TaskStatus = iota

	// Uncompleted representa una tarea que aún está pendiente
	Uncompleted
)

// String implementa la interfaz Stringer de Go, permitiendo imprimir TaskStatus de forma legible.
// Cuando se intenta convertir TaskStatus a string (ej: fmt.Println), este método se ejecuta automáticamente.
// Retorna:
// - "Completed" si el estado es 0
// - "Uncompleted" si el estado es 1
func (s TaskStatus) String() string {
	return [...] string{"Completed", "Uncompleted"}[s]
}

// IsCompleted es un método de utilidad que verifica si una tarea está completada.
// Ejemplo de uso: if task.Status.IsCompleted() { ... }
// Retorna true si el estado es Completed, false en caso contrario
func (s TaskStatus) IsCompleted() bool {
	return s == Completed
}
