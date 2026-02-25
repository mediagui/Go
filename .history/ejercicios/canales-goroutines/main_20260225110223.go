// Crear un programa en Go que realice lo siguiente:

// v Lanzar 3 goroutines.

// v Cada una deberá calcular la suma de un rango de números.

// v Enviar el resultado por un canal.

// v El programa principal deberá recoger los tres resultados y sumarlos.

package main

import (
	"crypto/rand"
	"log"
	"runtime"
	"sync"
)

func main() {

	var total int
	numOfGoRoutines := runtime.NumCPU()

	wg := sync.WaitGroup{}
	wg.Add(numOfGoRoutines)

	ch := make(chan int)
	defer close(ch)

	for range numOfGoRoutines {

		rand.Int()

		go emiteSuma([]int{1, 2, 3, 4, 5}, ch, &wg)
		total += recibeSuma(ch)
	}

	wg.Wait()

	println(total)

}

func emiteSuma(numeros []int, c chan<- int, wg *sync.WaitGroup) {

	suma := 0
	for _, n := range numeros {
		suma += n
	}

	log.Printf("Sumando valores del rango %v: %d\n", numeros, suma)

	wg.Done()

	c <- suma
}

func recibeSuma(ch <-chan int) int {
	return <-ch
}
