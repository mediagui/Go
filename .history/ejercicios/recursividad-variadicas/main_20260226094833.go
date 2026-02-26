// Se solicita crear un programa en Go, que realice lo siguiente:
// 1.    Función recursiva:
// Crear una función recursiva llamada [factorial] que reciba un número entero y devuelva su factorial.
// 2.    Función anónima:
// Dentro de [main], crear una función anónima que reciba un número entero y devuelva su cuadreado. Llamarla al menos dos veces
// 3.    Función varíadica:
// Crear una función [sumarTodo] que acepte un número variable de enteros y devuelva la suma de todos ellos. Probar con números distintos.

package main

func main() {}

func factorial(numero int) int {
	if numero == 0 {
		return 1
	}
	// q: Comparar con el rango maximo de un int
	// a: El rango máximo de un int en Go es 2^63 - 1 (para int64) o 2^31 - 1 (para int32), dependiendo de la arquitectura del sistema.
	// q: ¿como preguntamos por la arquitectura y comparamos con los dos tipos?
	// a:

	return numero * factorial(numero-1)
}
