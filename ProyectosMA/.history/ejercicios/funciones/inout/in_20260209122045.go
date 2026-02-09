package inout

import "funciones/dto"

type Dato struct{
	id int
	desc string
}

// Tipo pseudo-enumerado
type DatoTipo int

const (
	Nombre DatoTipo = iota
	Apellido1
	Apellido2
	Dni
	Tfno
)

var datoAPedir map[string]bool

func PedirInfo(user *dto.UsuarioStruct) {

}
