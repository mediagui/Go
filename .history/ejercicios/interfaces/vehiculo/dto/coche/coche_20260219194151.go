package coche

import "vehiculo/dto"

type Coche struct {
	dto.Vehiculo
	Matricula string
	Plazas    int
}
