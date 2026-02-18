// Package domain contiene las entidades principales del dominio de negocio.
// En la arquitectura hexagonal, el dominio es el núcleo más importante y debe
// ser INDEPENDIENTE de cualquier detalle técnico, adaptadores o frameworks.
// Esto significa:
// - No importa ningún otro paquete del proyecto
// - No depende de bases de datos específicas
// - No depende de tecnologías de transporte (HTTP, gRPC, etc)
package domain

// task es un struct PRIVADO que encapsula los datos internos de una tarea.
// El uso de minúscula inicial indica que no puede ser importado desde otros paquetes.
// Esto fuerza el uso de la interfaz pública TaskType que es el contrato externo.
// Ventajas de esta encapsulación:
// 1. Permite cambiar la implementación interna sin afectar usuarios del paquete
// 2. P fuerza el uso de constructores o métodos dedicados
// 3. Facilita la validación de datos
type task struct {
	// Id es el identificador único de la tarea
	Id uint8

	// Title es el título o nombre descriptivo de la tarea
	Title string

	// Description contiene detalles adicionales sobre la tarea
	Description string

	// Status representa el estado actual de la tarea (Completed o Uncompleted)
	Status TaskStatus
}

// TaskType es el tipo público que expone la tarea hacia fuera del paquete.
// Es un alias del struct privado 'task'.
// Este patrón permite:
// - Cambiar la implementación interna sin romper la interfaz pública
// - Mantener la compatibilidad hacia atrás
// - Encapsular detalles de implementación
//
// Usuarios del paquete ven e interactúan con TaskType, no con task
type TaskType task
