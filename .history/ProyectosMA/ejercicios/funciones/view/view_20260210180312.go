package view

import "fmt"

func MuestraMenuPrincipal() {
	var opcion int
	fmt.Println("1. Introduzca datos\n2. Seleccione campos a mostrar\n\tElija: ")
	fmt.Scanln(&opcion)
}
