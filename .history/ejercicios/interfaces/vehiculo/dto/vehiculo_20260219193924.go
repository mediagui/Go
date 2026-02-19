// v Tipo (coche, camión, moto, etc…)
// v Marca.
// v Modelo.
// v Matrícula.
// v Año del modelo.

package dto

type TipoVehiculo int

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

type Coche struct {
	Vehiculo
	Matricula string
	Plazas    int
}

type Patinete struct {
	Vehiculo
	Plazas int
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

func (a *Vehiculo) Arrancar() {

}

// v Tipo (coche, camión, moto, etc…)
// v Marca.
// v Modelo.
// v Matrícula.
// v Año del modelo.
