package moto

import "ejercicios/interfaces/vehiculo/dto/vehiculo"

type Moto struct {
	vehiculo.Vehiculo
	Matricula string
	Plazas    int
}
