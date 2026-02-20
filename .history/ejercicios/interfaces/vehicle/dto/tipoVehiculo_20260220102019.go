// Package dto provides data transfer objects for vehicle management.
package dto

// TipoVehiculo represents the type of a vehicle (car, truck, motorcycle, etc.).
type TipoVehiculo int

// String returns the string representation of the TipoVehiculo.
// It implements the fmt.Stringer interface.
func (t TipoVehiculo) String() string {
	switch t {
	case Car:
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
