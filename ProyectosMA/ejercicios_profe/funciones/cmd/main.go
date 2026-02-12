// Paquete principal al que pertenece el programa.
package main

import (
	"fmt"
	"funciones"
)

// Zona de declaración de variables.
var (
	// Variables para almacenar los códigos de colores.
	codigoRojo   = "\033[38;2;255;0;0m"
	codigoVerde  = "\033[38;2;0;255;0m"
	codigoAzul   = "\033[38;2;0;0;255m"
	codigoBlanco = "\033[38;2;255;255;255m"
	fondoBlanco  = "\033[48;2;255;255;255m"

	// Variable para almacenar el reset de los colores.
	reset = "\033[0m"
)

// Función principal.
func main() {
	fmt.Println("Saludos desde el main de funciones.")
	funciones.ExplicacionFunciones()
}
