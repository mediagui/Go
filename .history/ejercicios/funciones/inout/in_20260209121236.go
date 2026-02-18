package inout

import "funciones/dto"

// Tipo pseudo-enumerado
const (
	Nombre = iota
	Apellido1
	Apellido2
	Dni
	Tfno
)

var datoAPedir map[string]bool

func PedirInfo(user *dto.UsuarioStruct) {

}
