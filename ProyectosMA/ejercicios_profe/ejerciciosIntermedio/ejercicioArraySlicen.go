package ejerciciosIntermedio

import "fmt"

/*
###########
#ENUNCIADO#
###########

Crea 1 array números y dos Slicen. El primero tendrá 5 números, el segundo se llamará pares y el tercero impares, ambos estarán vacíos. Después multiplica cada uno de los números del primer array por un número introducido por el usuario entre 1 y 10, si el resultado es par guarda ese número en el Slicen de pares y si es impar en el Slicen de impares.

Muestra por consola la multiplicación que se produce junto con su resultado, con el formato:

2 x 3 = 6 -el Slicen de pares e impares.

*/

/*
###########
#VARIABLES#
###########
*/
var (
	// Variables para almacenar los códigos de colores.
	codigoRojo   = "\033[38;2;255;0;0m"
	codigoVerde  = "\033[38;2;0;255;0m"
	codigoAzul   = "\033[38;2;0;0;255m"
	codigoBlanco = "\033[38;2;255;255;255m"
	fondoBlanco  = "\033[48;2;255;255;255m"
	fondoNegro   = "\033[48;2;0;0;0m"

	// Variable para almacenar el reset de los colores.
	reset     = "\033[0m"
	resultado int
	numero    int
)

// Función que ejecuta el ejercicio. Utilizo los colores que se definen en el paquete main.
func EjercicioArraySlicen1() {
	fmt.Println(codigoAzul + fondoBlanco + `
######################################
#Ejecutando Ejercicio Arrays y Slices#
######################################
` + reset)

	// Creo un array con 5 número enteros.
	numeros := [5]int{}

	// Recorro el array "vacío" de numeros para almacenar los números que quiera el usuario para operar con ellos.
	for indice, _ := range numeros {
		fmt.Print("Indique el número para el índice [", indice, "]: ")
		fmt.Scan(&numeros[indice])
	}

	// Creo dos Slicen vacíos.
	pares := []int{}
	impares := []int{}

	// Creo un bucle for, para pedir al usuario 5 números en introducirlos en el array correspondiente.
	for i := 0; i < 5; i++ {
		// Solicito al usuario que introduzca un número entre el 1 y el 10 y lo almaceno en una variable.
		fmt.Print(i+1, " - Introduce un número entre 1 y 10: ")
		fmt.Scan(&numero)
		// Compruebo que el número esté dentro del rango y de no ser así, vuelvo a pedirlo.
		for numero < 1 || numero > 10 {
			// Muestro un mensaje de error y vuelvo a pedir el número.
			fmt.Println(codigoRojo + "Error: El número introducido no es válido." + reset)
			fmt.Print(i+1, " - Introduce un número entre 1 y 10: ")
			fmt.Scan(&numero)
		}
		// Multiplico el número introducido por el usuario, por el número correspondiente del arry.
		resultado = numero * numeros[i]

		// Muestro la multiplicación que se produce junto con su resultado.
		fmt.Println(numeros[i], " X ", numero, " = ", resultado)

		// Compruebo si el resultado es par o impar y lo almaceno en el Slicen correspondiente.
		if resultado%2 == 0 {
			// Si el resultado es par, lo almaceno en el slicen de pares.
			pares = append(pares, resultado)
		} else {
			// Si el resultado es impar, lo almaceno en el slicen de pares.
			impares = append(impares, resultado)
		}
	}

	// Muestro el slicen de pares e impares.
	if len(pares) > 0 && len(impares) > 0 {
		fmt.Println("Pares -> ", pares)
		fmt.Println("Impares -> ", impares)
	} else if len(pares) > 0 {
		fmt.Println("Pares -> ", pares)
	} else {
		fmt.Println("Impares -> ", impares)
	}
}
