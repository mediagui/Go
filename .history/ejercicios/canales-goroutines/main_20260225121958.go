// Crear un programa en Go que realice lo siguiente:
// v Lanzar 3 goroutines.
// v Cada una deberá calcular la suma de un rango de números.
// v Enviar el resultado por un canal.
// v El programa principal deberá recoger los tres resultados y sumarlos.

package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {

	numeroGoroutines := runtime.NumCPU() * 1000

	// Canal para recibir los resultados de las goroutines
	resultados := make(chan int, numeroGoroutines)

	// Lanzar 3 goroutines
	for i := 1; i <= numeroGoroutines; i++ {
		go calculaSumaEscribeEnCanal(i, resultados)
	}

	// Recoger los tres resultados y sumarlos
	totalSuma := calculaSumaTotalLeyendoDelCanal(numeroGoroutines, resultados)

	fmt.Printf("\nSuma total de los %d resultados: %d\n", numeroGoroutines, totalSuma)

}

func calculaSumaTotalLeyendoDelCanal(numeroGoruintas int, resultados chan int) int {
	var suma int
	for i := 0; i < numeroGoruintas; i++ {
		resultado := <-resultados
		fmt.Printf("Resultado de goroutine %d: %d\n", i+1, resultado)
		suma += resultado
	}
	return suma
}

// calculaSumaEscribeEnCanal calcula la suma de un rango de números y la envía por el canal
func calculaSumaEscribeEnCanal(id int, ch chan<- int) {

	// Generar rango de números
	cantidadDeNumeros := rand.Intn(100)
	numeros := generaNumeros(cantidadDeNumeros)

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
func generaNumeros(elementos int) []int {
	numeros := make([]int, elementos)
	for i := range numeros {
		numeros[i] = rand.Intn(100)
	}
	return numeros
}
