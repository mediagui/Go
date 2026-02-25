package main

func main() {}

func suma(numeros []int) {
	suma := 0
	for _, n := range numeros {
		suma += n
	}
	return suma
}
