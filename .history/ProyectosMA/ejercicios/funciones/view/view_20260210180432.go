package view

import "fmt"

func MuestraMenuPrincipal() {
	fmt.Println("1. Introduzca datos\n2. Seleccione campos a mostrar\n\tElija: ")

}

func GetOpcionMenuPrincipal() int {
	var opcion int
	fmt.Scanln(&opcion)
}
