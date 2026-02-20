package motorcycle

import "interfaces/vehicle/dto"

type Moto struct {
	dto.Vehicle
	Plate  string
	Places int
}
