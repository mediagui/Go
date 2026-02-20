// Indicop el nomrbe del paquete para la gestión de vehículos.
package vehiculo

// Importaciones.
import "fmt"

// Defino una estructura para crear los vehículos.
type Vehiculo struct {
	Tipo      string // Almaceno el tipo (coche, moto o camión)
	Marca     string // Almaceno la marca del fabricante.
	Modelo    string // Almaceno el modelo específico.
	Matricula string // Almaceno la matríucula
	Ano       int    // Almaceno el año de fabricación
}

/*
Constructor NuevoVehiculo:
Como en Go no esxisten los constructores tradicionales, creo esta función para instanciar y devolver un puntero a una nueva estructura de vehiculo inicializada.
*/
func NuevoVehiculo(tipo, marca, modelo, matricula string, ano int) *Vehiculo {
	// Devuelvo la dirección de memoria de un nuevo objeto [Vehiculo] con los valores recibidos.
	return &Vehiculo{
		Tipo:      tipo,
		Marca:     marca,
		Modelo:    modelo,
		Matricula: matricula,
		Ano:       ano,
	}
}

/*
Zona de Métodos del Vehículo:
Acciones que puede realizar cada instancia de vehículo
*/

// Método Arrancar: Informa que el vehículo ha arrancado.
func (v *Vehiculo) Arrancar() {
	// Evalúo el tipo de vehículo para mostrar el mensaje correspondiente.
	switch v.Tipo {
	case "moto", "Moto":
		fmt.Printf("La %s %s %s está arrancando.\n", v.Tipo, v.Marca, v.Modelo)
	case "coche", "Coche", "camion", "Camion", "camión", "Camión":
		fmt.Printf("El %s %s %s está arrancando.\n", v.Tipo, v.Marca, v.Modelo)
	default:
		// Si el tipo no coincide con los conocidos, devuelvo un mensaje genérico.
		fmt.Printf("No existe el tipo %s\n", v.Tipo)
	}
}

// Método Frenar: Informa que el vehículo está reduciendo la velocidad.
func (v *Vehiculo) Frenar() {
	// Evalúo el tipo de vehículo para mostrar el mensaje correspondiente.
	switch v.Tipo {
	case "moto", "Moto":
		fmt.Printf("La %s %s %s está frenando.\n", v.Tipo, v.Marca, v.Modelo)
	case "coche", "Coche", "camion", "Camion", "camión", "Camión":
		fmt.Printf("El %s %s %s está frenando.\n", v.Tipo, v.Marca, v.Modelo)
	default:
		// Si el tipo no coincide con los conocidos, devuelvo un mensaje genérico.
		fmt.Printf("No existe el tipo %s\n", v.Tipo)
	}
}

// Método Claxon: Informa que el vehículo está reduciendo la velocidad.
func (v *Vehiculo) Claxon() {
	// Evalúo el tipo de vehículo para mostrar el mensaje correspondiente.
	switch v.Tipo {
	case "moto", "Moto":
		fmt.Printf("La %s %s %s está tocando el claxon.\n", v.Tipo, v.Marca, v.Modelo)
	case "coche", "Coche", "camion", "Camion", "camión", "Camión":
		fmt.Printf("El %s %s %s está tocando el claxon.\n", v.Tipo, v.Marca, v.Modelo)
	default:
		// Si el tipo no coincide con los conocidos, devuelvo un mensaje genérico.
		fmt.Printf("No existe el tipo %s\n", v.Tipo)
	}
}

// Método ImprimirInfo: Muestra en pantalla todos los detalles técnicos del vehículo.
func (v *Vehiculo) ImprimirInfo() {
	// Imprimo la información del "objeto"
	fmt.Printf("Tipo: %s\nMarca: %s\nModelo: %s\nMatrícula: %s\nAño %d\n",
		v.Tipo, v.Marca, v.Modelo, v.Matricula, v.Ano)
}
