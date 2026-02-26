// Paquete principal.
package main

// Importaciones
import "fmt"

// 1. Función recursiva: factorial.
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// 3. Función variádica: suma todos los números recibidos.
func sumarTodo(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

var (
	numFactorial int
	numAnonima   int
)

// Función principal.
func main() {
	// Solicito al usuario un número a introducir.
	fmt.Print("Indique el número para calcular el caftorial: ")

	fmt.Scan(&numFactorial)

	fmt.Println("El factorial de ", numFactorial, " es: ", factorial(numFactorial))

	// 2. Función anónima: calcula el cuadrado.
	cuadrado := func(x int) int {
		return x * x
	}

	// Solicito al usuario un número para calcular el cuadrado.
	fmt.Print("Indique un número para calcular su cuadrado: ")
	fmt.Scan(&numAnonima)

	// Imprimo el cuadreado con la función anónima.
	fmt.Println("El cuadrado de ", numAnonima, " es ", cuadrado(numAnonima))

	// Solicito al usuario otro número para calcular el cuadrado.
	fmt.Print("Indique otro número para calcular su cuadrado: ")
	fmt.Scan(&numAnonima)

	// Imprimo el cuadreado con la función anónima.
	fmt.Println("El cuadrado de ", numAnonima, " es ", cuadrado(numAnonima))

	// Utilizo varias veces la función variadica.
	fmt.Println("Suma variadica (1, 2, 3): ", sumarTodo(1, 2, 3))
	fmt.Println("Suma variadica (5, 25, 15, 35, 45): ", sumarTodo(5, 25, 15, 35, 45))
	fmt.Println("Suma variadica (1, 67, 42, 76, 796): ", sumarTodo(1, 67, 42, 76, 796))
}
