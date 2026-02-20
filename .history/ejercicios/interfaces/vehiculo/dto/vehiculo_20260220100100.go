// v Tipo (coche, camión, moto, etc…)
// v Marca.
// v Modelo.
// v Matrícula.
// v Año del modelo.

package dto

const (
	Automovil TipoVehiculo = iota + 1
	Camion
	Motocicleta
	Patin
)

type Vehiculo struct {
	Tipo   TipoVehiculo
	Marca  string
	Modelo string
	Anio   uint16
}

// Constructor de automóviles
func NewVehiculo(auto Vehiculo) *Vehiculo {
	return &Vehiculo{
		Tipo:   auto.Tipo,
		Marca:  auto.Marca,
		Modelo: auto.Modelo,
		Anio:   auto.Anio,
	}
}
