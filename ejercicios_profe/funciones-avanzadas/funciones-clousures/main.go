package main

import "fmt"

/*
##########
#CLOUSERS#
##########

Son funciones anónimas que se definen dentro de otra función y que pueden acceder y modificar variables definidas en el ámbito de la función externa. Esto significa que un clousure puede "recordar" el valor de las variables del ámbito de la función en el momento en que se definen y utilizarlas en cualquier momento posterior.
*/

func incrementador() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// Declaro una variable de tipo [int]
	numEntero := incrementador()
	// Llamo a la función con el clousure.
	fmt.Println(numEntero())
	fmt.Println(numEntero())
	fmt.Println(numEntero())
	fmt.Println(numEntero())
}
