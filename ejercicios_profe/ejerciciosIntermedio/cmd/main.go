// Defino el paquete principal para la ejecución del programa.
package main

// Inicio la sección de importación de las dependencias requeridas.
import (
	// Importo el paquete de lógica interna para los ejercicios intermedios.
	"ejerciciosIntermedio"
	// Incorporo el paquete estándar para el manejo de salida formateada.
	"fmt"
)

// Zona de declaración de variables.
var (
	// Variables para almacenar los códigos de colores.
	codigoRojo   = "\033[38;2;255;0;0m"
	codigoVerde  = "\033[38;2;0;255;0m"
	codigoAzul   = "\033[38;2;0;0;255m"
	codigoBlanco = "\033[38;2;255;255;255m"
	fondoBlanco  = "\033[48;2;255;255;255m"
	fondoNegro   = "\033[48;2;0;0;0m"

	// Variable para almacenar el reset de los colores.
	reset = "\033[0m"

	// Variable para almacenar la opción elegida por el usuario.
	opcion int
	// Variable para salir del bucle.
	salir bool
)

// Establezco el punto de entrada principal para la ejecución de la aplicación.
func main() {

	// Utilizo un bucle infinito para imprimir el menú de ejercicios y ejecutar el ejercicio correspondiente.
	for !salir {
		// Verifico el inicio de la ejecución mediante un mensaje en la salida estándar.
		fmt.Print("\n" +
			fondoBlanco + codigoAzul + `
######
#MENÚ#
######
` + reset + `
` + codigoVerde + `1. Ejercicio Funciones` + reset + `
` + codigoVerde + `2. Ejercicio Array y Slicen` + reset + `
` + codigoVerde + `3. Ejercicio Mapas` + reset + `
` + codigoVerde + `4. Ejercicio Structs` + reset + `
` + codigoVerde + `5. Ejercicio Gestor de Tareas` + reset + `
` + codigoVerde + `6. Ejercicio Manejo de Errores` + reset + `
` + codigoRojo + `7. Salir` + reset + `
Seleccione una opción → `)
		// Capturo la opción elegida por el usuario.
		fmt.Scan(&opcion)

		// Utilizo una estructura [switch-case] para ejecutar la función del ejercicio correspondiente.
		switch opcion {
		case 1:
			ejerciciosIntermedio.EjercicioFunciones()
		case 2:
			ejerciciosIntermedio.EjercicioArraySlicen1()
		case 3:
			ejerciciosIntermedio.EjercicioMapas()
		case 4:
			ejerciciosIntermedio.CrearPersonitas()
		case 5:
			ejerciciosIntermedio.EjercicioGestorTareas()
		case 6:
			ejerciciosIntermedio.ManejoErrores()
		case 7:
			fmt.Println(codigoVerde + fondoNegro + `
#######################################
#Fin del programa. Gracias por usarlo.#
#######################################
` + reset)
			salir = true
		default:
			fmt.Println(codigoRojo + fondoNegro + "Opción no válida" + reset)
		}
	}
}
