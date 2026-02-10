// Paquete encargado de agrupar las funciones dedicadas a mostrar y recoger información de la consola
package view

import (
	"fmt"
	"funciones/dto"
	"slices"
	"strings"
)

// Muestra el menú principal para elegir entre introducir o mostrar datos
func MuestraMenuPrincipal() {
	fmt.Println("1. Introduzca datos\n2. Seleccione campos a mostrar\n\tElija: ")
}

func MuestraMenuVisualizacionDatos() {

}

// Lee en consola la opción seleccionada y la devuelve
func GetOpcionSeleccionada() int {
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}

// Lee por consola los campos a mostrar y devuelve la información almacenada en los mismos.
func GetCamposAMostrar() []string {

	var campos string
	fmt.Sscanln(campos)
	sliceCampos := strings.Split(strings.TrimSpace(campos), ",")

	// Eliminar los espacios en blanco de cada uno de los elementos del sliceCampos

	var i int
	for i < len(sliceCampos) {
		strings.TrimSpace(sliceCampos[i : i+1])
		i++
	}

	for i := 0; i < len(sliceCampos); i++ {
		item = strings.TrimSpace(sliceCampos[i])

	}

}
