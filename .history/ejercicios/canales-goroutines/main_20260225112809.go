// Crear un programa en Go que realice lo siguiente:
// v Lanzar 3 goroutines.
// v Cada una deberá calcular la suma de un rango de números.
// v Enviar el resultado por un canal.
// v El programa principal deberá recoger los tres resultados y sumarlos.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Canal para recibir los resultados de las goroutines
	resultados := make(chan int, 3)

	// Lanzar 3 goroutines
	for i := 1; i <= 3; i++ {
		go calcularSuma(i, resultados)
	}

	// Recoger los tres resultados y sumarlos
	totalSuma := 0
	for i := 0; i < 3; i++ {
		resultado := <-resultados
		fmt.Printf("Resultado de goroutine %d: %d\n", i+1, resultado)
		totalSuma += resultado
	}

	fmt.Printf("\nSuma total de los 3 resultados: %d\n", totalSuma)
}

// calcularSuma calcula la suma de un rango de números y la envía por el canal
func calcularSuma(id int, ch chan<- int) {
	// Generar rango de números
	numeros := generaNumeros()

	// Calcular la suma
	suma := 0
	for _, n := range numeros {
		suma += n
	}

	fmt.Printf("Goroutine %d - Rango: %v - Suma: %d\n", id, numeros, suma)

	// Enviar resultado por el canal
	ch <- suma
}

// generaNumeros genera un slice de números aleatorios
func generaNumeros() []int {
	numeros := make([]int, 5)
	for i := range numeros {
		numeros[i] = rand.Intn(100)
	}
	return numeros
}
