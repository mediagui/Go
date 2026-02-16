// Indico el paquete principal.
package main

// Importaciones.
import (
	"fmt"
	"manejoErrores"
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

// Declaro la función principal [main], que actúa como punto de inicio de la ejecución del programa. No recibe parámetros de entrada ni retorna ningún valor
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
` + codigoVerde + `1. Manejo de errores` + reset + `
` + codigoVerde + `2. Errores Personalizados` + reset + `
` + codigoVerde + `3. Ejemplo Manejo de Erres` + reset + `
` + codigoVerde + `4. Defer` + reset + `
` + codigoVerde + `5. Ejemplo Defer` + reset + `
` + codigoVerde + `6. Explicación Panic` + reset + `
` + codigoVerde + `7. Explicación Logs` + reset + `
` + codigoRojo + `8. Salir` + reset + `
Seleccione una opción → `)
		// Capturo la opción elegida por el usuario.
		fmt.Scan(&opcion)

		// Utilizo una estructura [switch-case] para ejecutar la función del ejercicio correspondiente.
		switch opcion {
		case 1:
			manejoErrores.ErroresBasicos()
		case 2:
			manejoErrores.ErrCustom()
		case 3:
			manejoErrores.EjpManejoErr()
		case 4:
			manejoErrores.ExplDefer()
		case 5:
			manejoErrores.EjpDefer()
		case 6:
			manejoErrores.ExplPanic()
		case 7:
			manejoErrores.RegistarErrores()
		case 8:
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
