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
	"vehiculo/dto"
)

func main() {

	vehiculos := buildVehicles()

	for _, v := range vehiculos {

		log.Printf("Vehicle: %v", v)

	}

}

func buildVehicles() map[string]dto.Vehicle {
	vehiculos := make(map[string]dto.Vehicle)

	coche := dto.NewVehiculo(dto.Vehicle{
		Tipo:   dto.Automovil,
		Marca:  "Toyota",
		Modelo: "Corolla",
		Anio:   uint16(2020),
	})

	camion := dto.NewVehiculo(dto.Vehicle{
		Tipo:   dto.Camion,
		Marca:  "Pegaso",
		Modelo: "Troner",
		Anio:   uint16(1985),
	})

	moto := dto.NewVehiculo(dto.Vehicle{
		Tipo:   dto.Motocicleta,
		Marca:  "Honda",
		Modelo: "CBR",
		Anio:   uint16(1990),
	})

	vehiculos[coche.Tipo.String()] = *coche
	log.Println("Adding a", coche.Tipo.String())
	vehiculos[camion.Tipo.String()] = *camion
	log.Println("Adding a", camion.Tipo.String())
	vehiculos[moto.Tipo.String()] = *moto
	log.Println("Adding a", moto.Tipo.String())

	return vehiculos
}
