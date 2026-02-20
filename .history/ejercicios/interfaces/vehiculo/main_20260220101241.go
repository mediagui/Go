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

import "vehiculo/dto"

func main() {

	vehiculos := buildVehiculos()
	println("Vehiculos creados:", len(vehiculos))
}

func buildVehiculos() map[string]dto.Vehiculo {
	vehiculos := make(map[string]dto.Vehiculo)

	coche := dto.NewVehiculo(dto.Vehiculo{
		Tipo:   dto.Automovil,
		Marca:  "Toyota",
		Modelo: "Corolla",
		Anio:   uint16(2020),
	})


	camion := dto.NewVehiculo(dto.Vehiculo{
		Tipo:   dto.Camion,
		Marca:  "Pegaso",
		Modelo: "Troner",
		Anio:   uint16(1985),
	})


	moto := dto.NewVehiculo(dto.Vehiculo{
		Tipo:   dto.Motocicleta,
		Marca:  "Honda",
		Modelo: "CBR",
		Anio:   uint16(1990),
	})


	vehiculos[coche.Tipo.String()] = *coche
	vehiculos[camion.Tipo.String()] = *camion
	vehiculos[moto.Tipo.String()] = *moto

	return vehiculos
}
