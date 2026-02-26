// Se solicita crear un programa en Go, que realice lo siguiente:
// 1.    Función recursiva:
// Crear una función recursiva llamada [factorial] que reciba un número entero y devuelva su factorial.
// 2.    Función anónima:
// Dentro de [main], crear una función anónima que reciba un número entero y devuelva su cuadreado. Llamarla al menos dos veces
// 3.    Función varíadica:
// Crear una función [sumarTodo] que acepte un número variable de enteros y devuelva la suma de todos ellos. Probar con números distintos.

package main

import "recursividad/internal"

func main() {

	internal.CalculoDeFactorial()

	internal.UsoDeFuncAnonima()

	internal.UsoDeFuncionVariadicaYAnonima()

}
