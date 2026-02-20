package dto

func (v Vehiculo) String() string {
	return v.Tipo.String()
}

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
