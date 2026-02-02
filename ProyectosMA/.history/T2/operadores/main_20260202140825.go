package main

import "fmt"

func main() {

	num := 2
	numCuad := num << 1
	fmt.Printf("Elevamos al cuadrado %v -> %v\n", num, numCuad)

	num = 36
	numSqr := (num >> 1) - 1
	fmt.Printf("Calculamos la raiz cuadrada %v -> %v\n", num, numSqr)

}
