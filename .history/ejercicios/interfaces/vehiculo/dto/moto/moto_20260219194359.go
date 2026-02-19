package moto

import "vehiculo/dto"

type Moto struct {
	dto.Vehiculo
	Matricula string
	Plazas    int
}
