// Package dto provides data transfer objects for vehicle management.
package dto

import (
	"fmt"
	"log"
)

// Vehicle type constants define the different types of vehicles supported.
const (
	// Car represents a car/automobile vehicle type.
	Car VehicleType = iota + 1
	// Lorry represents a truck vehicle type.
	Lorry
	// Motorcycle represents a motorcycle vehicle type.
	Motorcycle
	// Skateboard represents a scooter/skate vehicle type.
	Skateboard
)

// Vehicle represents a vehicle with its type, brand, model, and year.
type Vehicle struct {
	// Type is the type of vehicle (car, truck, motorcycle, etc.)
	Type VehicleType
	// Brand is the brand/make of the vehicle
	Brand string
	// Model is the model name of the vehicle
	Model string
	// Year is the year of the vehicle model
	Year uint16
	// Vehicle identification
	Plate string
}

func (v *Vehicle) Arrancar() {
	log.Println("Arrancando el vehículo", v.Plate)
}

// String returns the string representation of the Vehicle.
// It implements the fmt.Stringer interface by delegating to VehicleType.String().
func (v Vehicle) String() string {
	return fmt.Sprintf("[%s]: %s %s %s %d", v.Type.String(), v.Brand, v.Model, v.Plate, v.Year)
}

// NewVehicle creates and returns a new Vehicle pointer from the provided Vehicle value.
// This constructor creates a copy of the vehicle data on the heap.
func NewVehicle(auto Vehicle) *Vehicle {
	log.Println("Building a new", auto.Type.String())

	return &Vehicle{
		Type:  auto.Type,
		Brand: auto.Brand,
		Model: auto.Model,
		Year:  auto.Year,
		Plate: auto.Plate,
	}
}
