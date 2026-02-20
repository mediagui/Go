package moto

import "vehiculo/dto"

type Moto struct {
	dto.Vehicle
	Matricula string
	Plazas    int
}
