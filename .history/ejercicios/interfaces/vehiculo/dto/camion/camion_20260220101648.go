package camion

import "vehiculo/dto"

type Trailer struct {
	dto.Vehicle
	Matricula string
	Plazas    int
	Carga     uint
}
