package main

import (
	"ProyectosMA/variables"
	"fmt"
	"reflect"
)

func main() {
	variable := variables.GetUint8Value()
	fmt.Println("variable ", variable)
	fmt.Println("type:", reflect.TypeOf(variable))
}
