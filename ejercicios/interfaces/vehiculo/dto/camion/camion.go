package camion

import "vehiculo/dto"

type Trailer struct {
	dto.Vehiculo
	Matricula string
	Plazas    int
	Carga     uint
}
