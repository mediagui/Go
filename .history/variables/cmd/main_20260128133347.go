package main

import (
	"ProyectosMA/variables"
	"fmt"
)

func main() {
	variable := variables.GetUint8Value()
	fmt.Println("variable ", variable)
}
