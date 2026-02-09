package inout

import "funciones/dto"

type Dato struct{
	id int
	desc string
}

// Tipo pseudo-enumerado
const (
	Nombre Dato = {
		iota,
		Nombre
	}
	Apellido1
	Apellido2
	Dni
	Tfno
)

var datoAPedir map[string]bool

func PedirInfo(user *dto.UsuarioStruct) {

}
