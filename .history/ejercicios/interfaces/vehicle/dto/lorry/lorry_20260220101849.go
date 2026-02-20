package lorry

import "vehicle/dto"

type Trailer struct {
	dto.Vehicle
	Matricula string
	Plazas    int
	Carga     uint
}
