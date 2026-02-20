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

func (v Vehiculo) String() string {
	return v.Tipo.String()
}

func (t TipoVehiculo) String() string {
	switch t {
	case Automovil:
		return "Automovil"
	case Camion:
		return "Camion"
	case Motocicleta:
		return "Motocicleta"
	case Patin:
		return "Patín"
	default:
		panic("Tipo de vehiculo no reconocido")
	}
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
