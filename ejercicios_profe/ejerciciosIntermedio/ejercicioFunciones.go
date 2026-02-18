package ejerciciosIntermedio

import (
	"fmt"
)

/*
###########
#ENUNCIADO#
###########

Crear un programa en go, que solicite distintos datos a un usuario desde la consola (nombre, apellido 1, apellido 2, dni y nº tfno). El programa deberá mostrar un menú solicitando si desea imprimir todos los datos o solo alguno en concreto. Se deberán utilizar las funciones correspondientes para ello.
*/

/*
###########
#VARIABLES#
###########
*/
var (
	nombre    string
	apellido1 string
	apellido2 string
	dni       string
	tfno      string
	// Variable para almacenar la opción elegida por el usuario.
	opcion int
	// Variable para salir del bucle.
	salir bool
)

// Función de ejemplo para la plantilla
func EjercicioFunciones() {
	// Solicito al usuario que introduzca sus datos.
	fmt.Print("Introduce tu nombre: ")
	fmt.Scan(&nombre)
	fmt.Print("Introduce tu primer apellido: ")
	fmt.Scan(&apellido1)
	fmt.Print("Introduce tu segundo apellido: ")
	fmt.Scan(&apellido2)
	fmt.Print("Introduce tu DNI: ")
	fmt.Scan(&dni)
	fmt.Print("Introdue tu número de teléfono: ")
	fmt.Scan(&tfno)

	// Utilizo un bucle para mostrar el menú de opciones.
	for !salir {
		fmt.Print(`
1. Imprimir todos los datos
2. Mostrar solo algunos datos
4. Salir
¿Qué desea hacer?`)
		fmt.Scan(&opcion)
		// Ejecuto la acción del menú.
		switch opcion {
		case 1: // Mostrar todos los datos.
			imprimirDatos(nombre, apellido1, apellido2, dni, tfno)
		case 2: // Mostrar datos concretos.
			// Pido al usuario que indique qué dato desea imprimir.
			fmt.Print(`
1. Nombre
2. Apellido 1
3. Apellido 2
4. DNI
5. Teléfono
6. Salir
¿Qué dato desea imprimir?`)
			fmt.Scan(&opcion)
			// Llamo a la función para imprimir el dato solicitado.
			datoConcreto := imprimirDatoConcreto(opcion)
			fmt.Println(datoConcreto)
		case 4: // Salir del menú
			salir = true
		}
	}
}

// Función para imprimir todos los datos. Esta función recibirá como parámetros los datos del usuario y los imprimairá. No devuelve ningún valor.
func imprimirDatos(fnombre string, fapellido1 string, fapellido2 string, fdni string, ftfno string) {
	fmt.Printf(`
* Nombre: %s
* Apellido 1: %s
* Apellido 2: %s
* DNI: %s
* Teléfono: %s
`, fnombre, fapellido1, fapellido2, fdni, ftfno)
}

// Función que recibe la opción del menú y devuelve el dato que el usuario desea imprimir.
func imprimirDatoConcreto(fopcion int) (datoSalida string) {
	// Devuelvo el dato en función de la opción introducida por el usuario.
	switch fopcion {
	case 1: // Devuelvo el nombre
		datoSalida = nombre
	case 2: // Devuelvo el Apellido 1
		datoSalida = apellido1
	case 3: // Devuelvo el Apellido 2
		datoSalida = apellido2
	case 4: // Devuelvo el DNI
		datoSalida = dni
	case 5: // Devuelvo el Teléfono
		datoSalida = tfno
	case 6: // Devuelvo el Salgo del bucle
		return
	}
	return
}
