// Paquete encargado de agrupar las funciones dedicadas a mostrar y recoger información de la consola
package view

import (
	"fmt"
	"strings"
)

// Muestra el menú principal para elegir entre introducir o mostrar datos
func MuestraMenuPrincipal() {
	fmt.Println("1. Introduzca datos\n2. Seleccione campos a mostrar\n\tElija: ")
}

// Muestra el menú para introducir los campos que se desean mostrar
func MuestraMenuVisualizacionDatos() {
	fmt.Println("Los posibles campos a mostrar son: \n Nombre, Apellidos, Tfno y Dni \n Escriba los campos que quiere mostrar: ")
}

// Lee en consola la opción seleccionada y la devuelve
func GetOpcionSeleccionada() int {
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}

// Lee por consola los campos a mostrar y devuelve un slice con el nombre de los mismos.
func GetCamposAMostrar() []string {

	var campos string
	fmt.Sscanln(campos)
	sliceCampos := strings.Split(strings.TrimSpace(campos), ",")

	// Eliminar los espacios en blanco de cada uno de los elementos del sliceCampos
	for i := range sliceCampos {
		sliceCampos[i] = strings.TrimSpace(sliceCampos[i])
	}

	return sliceCampos

}
