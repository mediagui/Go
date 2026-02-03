package main

import (
	"fmt"
)

// Declaración de la función ensamblador
func Add(a, b int) int

func main() {
	x := 7
	y := 5
	r := Add(x, y)
	fmt.Println("Resultado:", r)
}
