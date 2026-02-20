package motorcycle

import "vehicle/dto"

type Moto struct {
	dto.Vehicle
	Matricula string
	Plazas    int
}
