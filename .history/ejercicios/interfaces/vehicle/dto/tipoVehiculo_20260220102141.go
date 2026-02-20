// Package dto provides data transfer objects for vehicle management.
package dto

// VehicleType represents the type of a vehicle (car, truck, motorcycle, etc.).
type VehicleType int

// String returns the string representation of the TipoVehiculo.
// It implements the fmt.Stringer interface.
func (t VehicleType) String() string {
	switch t {
	case Car:
		return "Car"
	case Lorry:
		return "Lorry"
	case Motorcycle:
		return "Motorcycle"
	case Skateboard:
		return "Skateboard"
	default:
		panic("Tipo de vehiculo no reconocido")
	}
}
