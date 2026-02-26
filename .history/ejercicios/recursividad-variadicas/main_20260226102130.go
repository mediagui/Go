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

	println("Calculamos el factorial de 10.000.")
	valor, err := factorial(10_000)
	if err != nil {
		println("Error calculando factorial:", err)
		return
	}
	println("Resultado:", valor)

	usoDeFuncAnonima()


}

// Calculo del factorial usando recursividad
func factorial(numero int) (int, error) {

	if numero == 0 {
		return 1, nil
	}

	if !internal.CanCalculateFactorial(numero) {
		panic("Número demasiado grande para calcular factorial en esta arquitectura")
	}

	result, err := factorial(numero - 1)
	if err != nil {
		return 0, err
	}

	return numero * result, nil
}

func usoDeFuncAnonima() {

	println("Usamos una funcion anónima para calcular el cuadrado de un número")

	// Función anónima que calcula el cuadrado de un número
	cuadrado := func(numero int) int {
		return numero * numero
	}

	// Llamada a la función anónima
	resultado1 := cuadrado(5)
	println("Cuadrado de 5:", resultado1)

}
