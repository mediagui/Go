package main

import "fmt"

func main() {

	leftShift(2)
	rightShift(9)

}

func leftShift(num int) {

	numCuad := num << 1
	fmt.Printf("Elevamos al cuadrado %v -> %v\n", num, numCuad)

}
func rightShift(num int) {

	numSqr := (num >> 1) - 1
	fmt.Printf("Calculamos la raiz cuadrada %v -> %v\n", num, numSqr)

}
