// Paquete principal.
package main

// Importaciones
import "fmt"

/*
######################
#FUNCIONES VARIADICAS#
######################

En go, las funciones variádicas son aquellas que pueden tomar un número variable de argumentos. Esto se logra utilizando la sintaxis [...tipo] después del tipo del último argumento de la función. El valor especifica el tipo de los argumentos variables que se pueden pasar a la función.
*/

// Función suma, que suma dos números y e imprime el resultado.
// func suma(numUno, numDos, numTres int) int { // Sin variádica.
// Le paso a la función suma 2 argumentos, uno fijo y el otro variable.
func suma(nombre string, nums ...int) int {
	var resultado int
	// resultado := numUno + numDos // Sin variádica.
	// fmt.Println("Función sumar") // Sin variádica.
	// fmt.Printf("%T - %v\n", nums, nums)
	for _, num := range nums {
		// Muestro la operación.
		fmt.Printf("Hola %s, la suma es: %d + %d = %d\n", nombre, resultado, num, resultado+num)
		resultado += num
	}
	return resultado
}

// Función que recibe distintos parámetros de distintos tipos.
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

func main() {
	fmt.Println("Funciones variadicas")
	// fmt.Println(suma(8, 56)) // Sin variádica.
	// resulSuma := suma(8, 96, 3435)
	// Recojo el resultado de la suma y lo almaceno en una variable.
	/* resulSuma := suma(8, 96, 3435)
	fmt.Println(resulSuma) */

	// Imprimo el resultado directamente de varias sumas.
	fmt.Println("Suma 1: ", suma("Diego", 2, 5, -17))
	fmt.Println("Suma 2: ", suma("Raul", 8, 35, 44, 20))
	fmt.Println("Suma 3: ", suma("Adolfo", 16, 22, 65, -989, 245, 55))

	// Llamo a la función para imprimir los distintos tipos de datos.
	imprimirDatos("Hola", 88, true, 3.14, "Adiós")
}
