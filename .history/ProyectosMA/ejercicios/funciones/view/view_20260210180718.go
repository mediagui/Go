// Paquete encargado de agrupar las funciones dedicadas a mostrar y recoger información de la consola
package view

import "fmt"

// Muestra el menú principal para elegir entre introducir o mostrar datos
func MuestraMenuPrincipal() {
	fmt.Println("1. Introduzca datos\n2. Seleccione campos a mostrar\n\tElija: ")
}

// Devuelve la opción seleccionada
func GetOpcionSeleccionada() int {
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}
