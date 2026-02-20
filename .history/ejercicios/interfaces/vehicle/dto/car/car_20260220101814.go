package coche

import "vehiculo/dto"

type Car struct {
	dto.Vehicle
	Matricula string
	Plazas    int
}
