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

	var v dto.Vehiculo

	vVehiculo := dto.NewVehiculo(dto.Vehiculo{
		Tipo:   dto.Automovil,
		Marca:  "Toyota",
		Modelo: "Corolla",
		Anio:   uint16(2020),
	})

}
