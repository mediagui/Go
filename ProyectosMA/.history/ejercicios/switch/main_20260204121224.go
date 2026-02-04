package main

import "fmt"

var num int

func main() {

	requestANumber(&num)

	number := getNumberTranscription(num)

	fmt.Println(number)

}

func getNumberTranscription(num int) string {
	var transcript string

	switch num {
	case 1:
		transcript = "One"
	case 2:
		transcript = "Two"
	case 3:
		transcript = "Three"
	case 4:
		transcript = "Four"
	case 5:
		transcript = "Five"
	default:
		transcript = "Unknown number"
	}

	return transcript
}

func requestANumber(num *int) {

	fmt.Print("Give me a number: ")
	fmt.Scanln(&num)

}
