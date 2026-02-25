package main

func main() {}

func suma(numeros []int, ch chan-> int) {
	suma := 0
	for _, n := range numeros {
		suma += n
	}
	return suma
}
