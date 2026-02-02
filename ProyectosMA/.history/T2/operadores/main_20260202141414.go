package main

import "fmt"

func main() {

	leftShift(2)
	rightShift(9)

}

func leftShift(num int) {

	numCuad := num << 1
	fmt.Printf("Desplazamos 1 bit a la izq %v -> %v\n", num, numCuad)

}
func rightShift(num int) {

	numSqr := (num >> 1) - 1
	fmt.Printf("Desplazamos 1 bit a la dcha %v -> %v\n", num, numSqr)

}

func isEven(num int) bool {
	return num & !num
}
