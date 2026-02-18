package main

import "fmt"

func main() {

	num := 2
	numCuad := num << 1
	fmt.Printf("Elevamos al cuadrado %v -> %v", num, numCuad)

	num = 9
	numSqr := num >> 1
	fmt.Printf("Calculamos la raiz cuadrada %v -> %v", num, numSqr)

}
