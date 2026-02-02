package main

import (
	"fmt"
	"paquetefmt"
)

func main() {
	paquetefmt.Formatfmt()

	var name string
	var age int
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)

}
