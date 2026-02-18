package main

import (
	"ProyectosMA/variables"
	"fmt"
	"reflect"
)

func main() {

	variable := variables.GetUint8Value()
	fmt.Println("variable type: ", reflect.TypeOf(variable), "value: ", variable)

}
