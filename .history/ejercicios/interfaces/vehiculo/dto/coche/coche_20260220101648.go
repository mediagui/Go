package coche

import "vehiculo/dto"

type Coche struct {
	dto.Vehicle
	Matricula string
	Plazas    int
}
