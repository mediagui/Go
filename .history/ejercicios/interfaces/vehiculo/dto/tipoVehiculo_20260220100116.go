package dto

type TipoVehiculo int

func (t TipoVehiculo) String() string {
	switch t {
	case Automovil:
		return "Automovil"
	case Camion:
		return "Camion"
	case Motocicleta:
		return "Motocicleta"
	case Patin:
		return "Patín"
	default:
		panic("Tipo de vehiculo no reconocido")
	}
}
