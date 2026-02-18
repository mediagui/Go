package main

import (
	"fmt"
	m "moduletoimport"
)

func Greeting(name string) string {
	return "Hola, " + name
}

func main() {

	fmt.Println(m.Greeting("Míguel"))

}
