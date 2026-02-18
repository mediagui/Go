package main

import "fmt"

var num int

func main() {

	var msg string

	if num == 0 {
		msg = "Es neutro"
	} else if result := isOdd(num); result {
		msg = "Es impar"
	} else {
		msg = "Es par"
	}

	fmt.Println(msg)
}

func requestData() {
	fmt.Print("Número: ")
	fmt.Scanln(&num)
}

func isOdd(num int) bool {
	return num<<1&1 == 1
}
