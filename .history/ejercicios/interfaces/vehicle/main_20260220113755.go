// Tema 3 Ejercicio 7: Creación de autos
// Crear un programa en go que permita gestionar un concesionario, creando automóviles.
// De cada auto, se deberán especificar:
// v Tipo (coche, camión, moto, etc…)
// v Marca.
// v Modelo.
// v Matrícula.
// v Año del modelo.

// Se deberán especificar las acciones para cada coche:
// v  Arrancar.
// v  Frenar.
// v  Claxon.
package main

import (
	"log"
	"vehicle/dto"
)

func main() {

	vehiculos := buildVehicles()

	printVehicleDetails(vehiculos)

	startEngines(vehiculos)

	stopEngines(vehiculos)

}

func startEngines(vehiculos map[string]dto.Vehicle) {
	for _, v := range vehiculos {
		v.StartEngine()
	}
}

func stopEngines(vehiculos map[string]dto.Vehicle) {
	for _, v := range vehiculos {
		v.StopEngine()
	}
}

func makeClaxonSound(vehiculos map[string]dto.Vehicle) {
	for _, v := range vehiculos {
		v.MakeSound()
	}
}

func printVehicleDetails(vehiculos map[string]dto.Vehicle) {
	for _, v := range vehiculos {

		log.Printf("Vehicle: %v", v)

	}
}

func buildVehicles() map[string]dto.Vehicle {
	vehiculos := make(map[string]dto.Vehicle)

	coche := dto.NewVehicle(dto.Vehicle{
		Type:  dto.Car,
		Brand: "Toyota",
		Model: "Corolla",
		Year:  uint16(2020),
		Plate: "4866-BRJ",
	})

	camion := dto.NewVehicle(dto.Vehicle{
		Type:  dto.Lorry,
		Brand: "Pegaso",
		Model: "Troner",
		Year:  uint16(1985),
		Plate: "1234-ABC",
	})

	moto := dto.NewVehicle(dto.Vehicle{
		Type:  dto.Motorcycle,
		Brand: "Honda",
		Model: "CBR",
		Year:  uint16(1990),
		Plate: "5678-DEF",
	})

	vehiculos[coche.Type.String()] = *coche
	log.Println("Adding a", coche.Type.String())
	vehiculos[camion.Type.String()] = *camion
	log.Println("Adding a", camion.Type.String())
	vehiculos[moto.Type.String()] = *moto
	log.Println("Adding a", moto.Type.String())

	return vehiculos
}
