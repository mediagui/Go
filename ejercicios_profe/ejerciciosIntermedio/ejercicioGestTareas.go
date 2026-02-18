// Defino el paquete en el que irá el prgroama.
package ejerciciosIntermedio

/*
# [bufio]
				-> Proporciona lectura de búfer, lo que nos permite leer textos con espacios cómodamente.
# [os]
				-> Nos permite interactuar con el sistema operativo, en este caso, para acceder al teclado [Stdin].
*/
import (
	"bufio"
	"fmt"
	"os"
)

/*
###########
#Enunciado#
###########

#Gestor de Tareas#

Crear una aplicación en go que permita gestionar una lista de tareas (CRUD básico).
Se debe poder:
- Agregar tareas.
- Marcar como completadas.
- Editar tareas.
- Eliminar tareas.
*/

// Defino la estructura básica de las tareas.
type Tarea struct {
	// Defino los atributos de la tarea.
	nombre      string
	descripcion string
	completado  bool
}

// Creo otra estructurta que contendrá la colección de tareas.
type ListaTareas struct {
	// Slice en el que almacenaré los objetos de tipo [Tarea]
	tareas []Tarea
}

// Método para añadir una nueva tarea a la lista.
func (m_lista *ListaTareas) agregarTarea(m_tarea Tarea) {
	// Utilizo [append()], para añadir cada tarea al final de la lista.
	m_lista.tareas = append(m_lista.tareas, m_tarea)
}

// Método para "marcar completada" una tarea.
func (m_lista *ListaTareas) marcarCompletado(m_index int) {
	// Accedo a la tarea en la posición indicada (ajustando a 0) y cambiao su estado a [true]
	m_lista.tareas[m_index-1].completado = true
}

// Método pàra "Editar" una tarea existente.
func (m_lista *ListaTareas) editarTarea(m_index int, m_tarea Tarea) {
	// Reemplazo la tarea antigua en esa posición, por la nueva versión "m_tarea"
	m_lista.tareas[m_index-1] = m_tarea
}

// Método para "Eliminar" una tarea.
func (m_lista *ListaTareas) eliminarTarea(m_index int) {
	// Para eliminar una tarea de la lista, "cortamos" el slice quitando el elemento en la posición "m_index-1"
	m_lista.tareas = append(m_lista.tareas[:m_index-1], m_lista.tareas[m_index:]...)
}

// Método par amostrar la lista de tareas de forma ordenada.
func (m_lista *ListaTareas) mostrarTareas() {
	// Imprimo el contenido de la lista.
	fmt.Println("_-=Listado de tareas actualizado=-_")
	fmt.Println("###########################################")
	// Recorro el slice de tareas para imprimir cada una.
	for i, tarea := range m_lista.tareas {
		// Evalúo si la tarea está terminada para mostrar un mensaje personalizado en consecuencia.
		if tarea.completado {
			// Si está completada, muestro el mensaje correspondiente.
			fmt.Printf("%d. %s - %s - Completado: Tarea completada correctamente (%t)\n", i+1, tarea.nombre, tarea.descripcion, tarea.completado)
		} else {
			// Si está pendiente, mostramos otro mensaje.
			fmt.Printf("%d. %s - %s - Completado: Tarea pendiente. (%t)\n", i+1, tarea.nombre, tarea.descripcion, tarea.completado)
		}
	}
	fmt.Println("###########################################")
}

// Función para gestionar las tareas.
func EjercicioGestorTareas() {
	// Creo una instancia vacía de la lista de tareas, para poder empezar a trabajar.
	lista := ListaTareas{}

	// Preparo el lector del buffer.
	consola := bufio.NewReader(os.Stdin)

	// Inicio el bucle para el menú.
	for {
		// Variable para almacenar la opción introducida por el usuario.
		var opcion int
		// Muestro el menú.
		fmt.Print(`
######
#MENÚ#
######

1. Mostrar tareas
2. Añadir tarea
3. Marcar tarea como completada
4. Editar tarea
5. Eliminar tarea
6. Salir

Escoja una opción -> `)
		// Leo la opción introducida por el usuario.
		fmt.Scanln(&opcion)

		// Utilizo [switch-case] para ejecutar la opción indicada por el usuario.
		switch opcion {
		case 1: // Mostrar lista de tareas.
			lista.mostrarTareas()
		case 2: // Añadir nueva tarea.
			var tarea Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			// Leeo el nombre usando el lector que he creado anteriormente.
			tarea.nombre, _ = consola.ReadString('\n')
			// Solicito la descripción.
			fmt.Println("Ingrese la descripción de la tarea:")
			tarea.descripcion, _ = consola.ReadString('\n')
			// Añado la tarea a la colección.
			lista.agregarTarea(tarea)
			fmt.Println("Tarea añadida correctamente.")
		case 3: // Completar tarea.
			// Solicito el índice de la tarea que queremos "completar."
			var index int
			fmt.Print("Ingrese el índice de la tarea a completar: ")
			fmt.Scanln(&index)
			// Llamo al método que cambia el booleano.
			lista.marcarCompletado(index)
			// Indico que se ha completado correctamente.
			fmt.Println("Tarea maracada como completada correctamente.")
		case 4: // Editar tarea
			// Creo variable para almacenar el índice e instancio una tarea.
			var index int
			var tarea Tarea
			// Solicito el la tarea a actualizar.
			fmt.Print("Ingrese el índice de la tarea que desea actualizar:")
			fmt.Scanln(&index)
			// Solicito los nuevos valores.
			fmt.Print("Ingrese el nuevo nombre de la tarea: ")
			tarea.nombre, _ = consola.ReadString('\n')
			fmt.Print("Ingrese la nueva descripción de la tarea: ")
			tarea.descripcion, _ = consola.ReadString('\n')
			// Sobrescribimos la tarea en esa posición.
			lista.editarTarea(index, tarea)
			// Indico que la tarea se ha actualizado correctamente.
			fmt.Println("Tarea actualizada correctamente.")
		case 5: // Borrar tarea.
			// solicito el índice de la tarea a eliminar.
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			// Llamo al método que me permite quitar la tarea de la lista.
			lista.eliminarTarea(index)
			// Indico que la tarea se ha eliminado correctamente.
			fmt.Println("Tarea eliminada correctamente.")
		case 6: // Salir del programa.
			fmt.Println("Saliendo del programa de gestión de tareas 😒...")
			// Salgo del menú.
			return
		default: // Opción incorrecta.
			fmt.Println("Opción, como clarita, inválida.")
		}
	}
}
