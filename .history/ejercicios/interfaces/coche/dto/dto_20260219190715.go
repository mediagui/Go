// v Tipo (coche, camión, moto, etc…)
// v Marca.
// v Modelo.
// v Matrícula.
// v Año del modelo.

package dto


type TipoVehiculo int

const (
	Coche TipoVehiculo = iota
	Camion
	Moto
)

type Auto struct {
	Tipo      TipoVehiculo
	Marca     string
	Modelo    string
	Matricula string
	Anio      int
}