// Package dto provides data transfer objects for vehicle management.
package dto

import "log"

// Vehicle type constants define the different types of vehicles supported.
const (
	// Car represents a car/automobile vehicle type.
	Car TipoVehiculo = iota + 1
	// Lorry represents a truck vehicle type.
	Lorry
	// Motocicleta represents a motorcycle vehicle type.
	Motocicleta
	// Patin represents a scooter/skate vehicle type.
	Patin
)

// Vehicle represents a vehicle with its type, brand, model, and year.
type Vehicle struct {
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
func (v Vehicle) String() string {
	return v.Tipo.String()
}

// NewVehiculo creates and returns a new Vehiculo pointer from the provided Vehiculo value.
// This constructor creates a copy of the vehicle data on the heap.
func NewVehiculo(auto Vehicle) *Vehicle {
	log.Println("Building a new vehicle. This is a", auto.Tipo.String())

	return &Vehicle{
		Tipo:   auto.Tipo,
		Marca:  auto.Marca,
		Modelo: auto.Modelo,
		Anio:   auto.Anio,
	}
}
