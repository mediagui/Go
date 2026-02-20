package car

import "vehicle/dto"

type Car struct {
	dto.Vehicle
	Matricula string
	Plazas    int
}
