package main

import (
	"log"
	"sync"
)

func main() {

	var total int

	wg := sync.WaitGroup{}
	wg.Add(3)

	ch := make(chan int)
	defer close(ch)

	for i := range 3 {
		go emiteSuma([]int{1, 2, 3, 4, 5}, ch, &wg)
	}

	wg.Wait()

	total += recibeSuma(ch)
	println(total)

}

func emiteSuma(numeros []int, c chan<- int, wg *sync.WaitGroup) {
	suma := 0
	for _, n := range numeros {
		suma += n
	}

	log.Printf("Calculando suma en goroutine: %d\n", suma)

	wg.Done()

	log.Println("Goroutine finalizada.")
	c <- suma
}
func recibeSuma(ch <-chan int) int {
	return <-ch
}
