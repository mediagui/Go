// Package dto provides data transfer objects for vehicle management.
package dto

import (
	"fmt"
	"log"

	"github.com/gen2brain/beeep"
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

type driveable interface{
	StartEngine()
	StopEngine()
	MakeSound()
}


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
	// Vehicle status
	EngineStarted bool
}

// Starts the vehicle engine
func (v *Vehicle) StartEngine() {
	log.Println("Starting engine")
	v.EngineStarted = true
	log.Println("Engine started for vehicle", v.Plate)

	log.Println("Brum, Brum!!")

}

// Stops the vehicle engine
func (v *Vehicle) StopEngine() {
	log.Println("Stopping engine")
	v.EngineStarted = false
	log.Println("Engine stopped for vehicle", v.Plate)
}

func (v Vehicle) MakeSound() {
	switch v.Type {
	case Lorry:
		// Truck: lower frequency and longer duration
		beeep.Beep(600, 600)
	case Car:
		beeep.Beep(800, 600)
	case Motorcycle:
		beeep.Beep(1000, 600)
	default:
		// Standard sound for other vehicles
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	}

	log.Println("A ", v.Type.String(), "is beeping")
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
