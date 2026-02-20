// Package dto provides data transfer objects for vehicle management.
package dto

import "log"

// Vehicle type constants define the different types of vehicles supported.
const (
	// Automovil represents a car/automobile vehicle type.
	Automovil TipoVehiculo = iota + 1
	// Camion represents a truck vehicle type.
	Camion
	// Motocicleta represents a motorcycle vehicle type.
	Motocicleta
	// Patin represents a scooter/skate vehicle type.
	Patin
)

// Vehiculo represents a vehicle with its type, brand, model, and year.
type Vehiculo struct {
	// Tipo is the type of vehicle (car, truck, motorcycle, etc.)
	Tipo TipoVehiculo
	// Marca is the brand/make of the vehicle
	Marca string
	// Modelo is the model name of the vehicle
	Modelo string
	// Anio is the year of the vehicle model
	Anio uint16
}

// String returns the string representation of the Vehiculo.
// It implements the fmt.Stringer interface by delegating to TipoVehiculo.String().
func (v Vehiculo) String() string {
	return v.Tipo.String()
}

// NewVehiculo creates and returns a new Vehiculo pointer from the provided Vehiculo value.
// This constructor creates a copy of the vehicle data on the heap.
func NewVehiculo(auto Vehiculo) *Vehiculo {
	log.Println("Building a new vehicle. This is a", auto.Tipo.String())
	return &Vehiculo{
		Tipo:   auto.Tipo,
		Marca:  auto.Marca,
		Modelo: auto.Modelo,
		Anio:   auto.Anio,
	}
}
