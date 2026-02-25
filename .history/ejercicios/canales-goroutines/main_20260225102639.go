package main

func main() {}

func emiteSuma(numeros []int, ch chan<- int) {
	suma := 0
	for _, n := range numeros {
		suma += n
	}

	ch <- suma
}
