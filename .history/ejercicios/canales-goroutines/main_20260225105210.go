package main

import (
	"log"
	"runtime"
	"sync"
)

func main() {

	var total int
	numOfGoRoutines:=runtime.NumCPU()

	wg := sync.WaitGroup{}
	wg.Add(numOfGoRoutines)


	ch := make(chan int)
	defer close(ch)

	for range numOfGoRoutines {
		go emiteSuma([]int{1, 2, 3, 4, 5}, ch, &wg, runtime.NumGoroutine())
		total += recibeSuma(ch)
	}

	wg.Wait()

	println(total)

}

func emiteSuma(numeros []int, c chan<- int, wg *sync.WaitGroup, goRoutineNumber int) {
	suma := 0
	for _, n := range numeros {
		suma += n
	}

	log.Printf("Calculando suma en goroutine {%d}: %d\n", goRoutineNumber, suma)

	wg.Done()

	log.Println("Goroutine finalizada.")
	c <- suma
}
func recibeSuma(ch <-chan int) int {
	return <-ch
}
