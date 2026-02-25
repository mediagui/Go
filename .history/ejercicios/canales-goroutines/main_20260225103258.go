package main

func main() {

	ch := make(chan int)
	defer close(ch)

	go emiteSuma([]int{1, 2, 3, 4, 5}, ch)
	go emiteSuma([]int{1, 2, 3, 4, 5}, ch)
	go emiteSuma([]int{1, 2, 3, 4, 5}, ch)

	suma := recibeSuma(ch)
	println(suma)

}

func emiteSuma(numeros []int, c chan<- int) {
	suma := 0
	for _, n := range numeros {
		suma += n
	}

	c <- suma
}
func recibeSuma(ch <-chan int) int {
	sumaRecibida := 0
	sumaRecibida += <-ch
	return <-ch
}
