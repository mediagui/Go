package main

import (
	"ProyectosMA/variables"
	"fmt"
	"reflect"
)

func main() {

	variable := variables.GetUint8Value()
	fmt.Println("var type:\t", reflect.TypeOf(variable), "\tvalue:\t", variable)

	variables.TratoVariables()

}
