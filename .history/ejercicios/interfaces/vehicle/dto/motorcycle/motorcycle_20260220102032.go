package motorcycle

import "vehicle/dto"

type Moto struct {
	dto.Vehicle
	Plate string
	Places    int
}
