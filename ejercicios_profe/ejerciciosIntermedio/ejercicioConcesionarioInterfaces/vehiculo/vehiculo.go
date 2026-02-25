package vehiculo

/*
###########
#ENUNCIADO#
###########

Tema 3 Ejercicio 10: Creación de autos (Interfaces)

Modificar el programa que permite gestionar un concesionario, creando automóviles, de manera que los métodos de arrancar, frenar y tocar el claxon se implementen utilizando interface.
*/

import (
	"fmt"
)

// Defino una estructura para crear los vehículos.
type Vehiculo struct {
	tipo      string // Almaceno el tipo (coche, moto o camión)
	marca     string // Almaceno la marca del fabricante.
	modelo    string // Almaceno el modelo específico.
	matricula string // Almaceno la matríucula
	ano       int    // Almaceno el año de fabricación
}

// Estructura para el coche, que va a heredar de Vehiculo.
type Coche struct {
	Vehiculo       // Campo "embebido", permite acceder directamente a los atributos de vehiculo
	puertas        int
	motor          string
	esDescapotable bool
}

// Estructura para la moto, que va a heredar de Vehiculo.
type Moto struct {
	Vehiculo   // Campo "embebido", permite acceder directamente a los atributos de vehiculo
	cilindrada int
	potencia   int
}

// Estructura para el camión, que va a heredar de Vehiculo.
type Camion struct {
	Vehiculo  // Campo "embebido", permite acceder directamente a los atributos de vehiculo
	mma       int
	tipoCarga string
}

/*
Constructor NuevoVehiculo:
Como en Go no esxisten los constructores tradicionales, creo esta función para instanciar y devolver un puntero a una nueva estructura de vehiculo inicializada.
*/
func NuevoVehiculo(tipo, marca, modelo, matricula string, ano int) *Vehiculo {
	// Devuelvo la dirección de memoria de un nuevo objeto [Vehiculo] con los valores recibidos.
	return &Vehiculo{
		tipo:      tipo,
		marca:     marca,
		modelo:    modelo,
		matricula: matricula,
		ano:       ano,
	}
}

// Constructor para el nuevo coche.
func NuevoCoche(tipo, marca, modelo, matricula string, ano, puertas int, motor string, esDescapotable bool) *Coche {
	return &Coche{
		// Inicializo la estructura de vehiculo.
		Vehiculo: Vehiculo{
			tipo:      tipo,
			marca:     marca,
			modelo:    modelo,
			matricula: matricula,
			ano:       ano,
		},
		puertas:        puertas,
		motor:          motor,
		esDescapotable: esDescapotable,
	}
}

/*
Zona de Métodos del Vehículo:
Acciones que puede realizar cada instancia de vehículo
*/

// Métodos getters y setters para tratar con los datos del vehículo.
func (v *Vehiculo) GetTipo() string {
	return v.tipo
}

func (v *Vehiculo) SetTipo(tipo string) error {
	if tipo == "" {
		return fmt.Errorf("el tipo no puede estar vacío.")
	}
	v.tipo = tipo
	return nil
}

func (v *Vehiculo) GetMarca() string {
	return v.marca
}

func (v *Vehiculo) SetMarca(marca string) error {
	if marca == "" {
		return fmt.Errorf("la marca no puede estar vacío.")
	}
	v.marca = marca
	return nil
}

func (v *Vehiculo) GetModelo() string {
	return v.modelo
}

func (v *Vehiculo) SetModelo(modelo string) error {
	if modelo == "" {
		return fmt.Errorf("el modelo no puede estar vacío.")
	}
	v.modelo = modelo
	return nil
}

func (v *Vehiculo) GetMatricula() string {
	return v.matricula
}

func (v *Vehiculo) SetMatricula(matricula string) error {
	if matricula == "" {
		return fmt.Errorf("la matricula no puede estar vacío.")
	}
	v.matricula = matricula
	return nil
}

func (v *Vehiculo) GetAno() int {
	return v.ano
}

func (v *Vehiculo) SetAno(ano int) error {
	if ano == 0 || ano < 1886 {
		return fmt.Errorf("el ano no es válido.")
	}
	v.ano = ano
	return nil
}

// Método Arrancar: Informa que el vehículo ha arrancado.
func (v *Vehiculo) Arrancar() {
	// Evalúo el tipo de vehículo para mostrar el mensaje correspondiente.
	switch v.tipo {
	case "moto", "Moto":
		fmt.Printf("La %s %s %s está arrancando.\n", v.tipo, v.marca, v.modelo)
	case "coche", "Coche", "camion", "Camion", "camión", "Camión":
		fmt.Printf("El %s %s %s está arrancando.\n", v.tipo, v.marca, v.modelo)
	default:
		// Si el tipo no coincide con los conocidos, devuelvo un mensaje genérico.
		fmt.Printf("No existe el tipo %s\n", v.tipo)
	}
}

// Método Frenar: Informa que el vehículo está reduciendo la velocidad.
func (v *Vehiculo) Frenar() {
	// Evalúo el tipo de vehículo para mostrar el mensaje correspondiente.
	switch v.tipo {
	case "moto", "Moto":
		fmt.Printf("La %s %s %s está frenando.\n", v.tipo, v.marca, v.modelo)
	case "coche", "Coche", "camion", "Camion", "camión", "Camión":
		fmt.Printf("El %s %s %s está frenando.\n", v.tipo, v.marca, v.modelo)
	default:
		// Si el tipo no coincide con los conocidos, devuelvo un mensaje genérico.
		fmt.Printf("No existe el tipo %s\n", v.tipo)
	}
}

// Método Claxon: Informa que el vehículo está reduciendo la velocidad.
func (v *Vehiculo) Claxon() {
	// Evalúo el tipo de vehículo para mostrar el mensaje correspondiente.
	switch v.tipo {
	case "moto", "Moto":
		fmt.Printf("La %s %s %s está tocando el claxon.\n", v.tipo, v.marca, v.modelo)
	case "coche", "Coche", "camion", "Camion", "camión", "Camión":
		fmt.Printf("El %s %s %s está tocando el claxon.\n", v.tipo, v.marca, v.modelo)
	default:
		// Si el tipo no coincide con los conocidos, devuelvo un mensaje genérico.
		fmt.Printf("No existe el tipo %s\n", v.tipo)
	}
}

// Método ImprimirInfo: Muestra en pantalla todos los detalles técnicos del vehículo.
func (v *Vehiculo) ImprimirInfo() {
	// Imprimo la información del "objeto"
	fmt.Printf("Tipo: %s\nMarca: %s\nModelo: %s\nMatrícula: %s\nAño %d\n",
		v.tipo, v.marca, v.modelo, v.matricula, v.ano)
}
