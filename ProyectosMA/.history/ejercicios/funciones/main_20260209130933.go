// Crear un programa en go, que solicite distintos datos a un usuario desde la
// consola (nombre, apellido 1, apellido 2, dni y nº tfno).
// El programa deberá mostrar un menú solicitando si desea imprimir todos los datos o solo alguno en concreto.
// Se deberán utilizar las funciones correspondientes para ello.

package main

import (
	"fmt"
	"funciones/dto"
	"funciones/inout"
)

var usuario *dto.UsuarioStruct

func main() {

	inout.PedirInfo(usuario)

	fmt.Printf("usuario: %#v", &usuario)
}
