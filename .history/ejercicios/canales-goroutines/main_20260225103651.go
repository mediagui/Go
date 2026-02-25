package main

import "sync"

func main() {

	wg := sync.WaitGroup{}
	wg.Add(3)

	ch := make(chan int)
	defer close(ch)

	go emiteSuma([]int{1, 2, 3, 4, 5}, ch, &wg)
	go emiteSuma([]int{1, 2, 3, 4, 5}, ch, &wg)
	go emiteSuma([]int{1, 2, 3, 4, 5}, ch, &wg)

	wg.Wait()

	suma := recibeSuma(ch)
	println(suma)

}

func emiteSuma(numeros []int, c chan<- int, wg *sync.WaitGroup) {
	suma := 0
	for _, n := range numeros {
		suma += n
	}

	wg.Done()

	c <- suma
}
func recibeSuma(ch <-chan int) int {
	sumaRecibida := 0
	sumaRecibida += <-ch
	return <-ch
}
