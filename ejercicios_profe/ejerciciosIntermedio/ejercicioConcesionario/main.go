// Definición del paquete principal.
package main

// Importaciones
import (
	"bufio" // Paquete para la lectura eficiente de datos de entrada.
	"concesionario/vehiculo"
	"fmt"     // Paquete estándar para entrada y salida formateada.
	"os"      // Paquete para interactuar con el sistema operativo (entrada estándar)
	"strconv" // Paquete para converión de cadenas a otros tipos (string a int)
	"strings" // Paquete para manipular cadenas de texto.
)

/*
###########
#ENUNCIADO#
###########

Crear un programa en go que permita gestionar un concesionario, creando automóviles.

De cada auto, se deberán especificar:

v Tipo (coche, camión, moto, etc…)

v Marca.

v Modelo.

v Matrícula.

v Año del modelo.

Se deberán especificar las acciones para cada coche:

v  Arrancar.

v  Frenar.

v  Claxon.
*/

// Variables.
var (
	tipo      string
	marca     string
	modelo    string
	matricula string
	ano       int
	opcion    int  // Almacena la selección del menú del usuario.
	salir     bool // Controla la salida del bucle del menú.
)

func main() {
	/*
		Capturo los datos iniciales:
		Preparo un lector para recoger información detallada del vehículo.
	*/

	// Instancio el lector.
	rd := bufio.NewReader(os.Stdin)

	// Solicito el tipo de vehículo.
	fmt.Print("Ingrese un tipo de vehículo: ")
	fmt.Scanln(&tipo)

	// Solicito la marca.
	fmt.Print("Ingrese la marca del vehículo: ")
	marca, _ = rd.ReadString('\n')
	marca = strings.TrimSpace(marca)

	// Solicito el modelo.
	fmt.Print("Ingrese el modelo del vehículo: ")
	modelo, _ = rd.ReadString('\n')
	modelo = strings.TrimSpace(modelo)

	// Solicito la matrícula.
	fmt.Print("Ingrese la matrícula del vehículo: ")
	matricula, _ = rd.ReadString('\n')
	matricula = strings.TrimSpace(matricula)

	// Solicito el año de fabricación.
	fmt.Print("Ingrese la año del vehículo: ")
	input, _ := rd.ReadString('\n')
	input = strings.TrimSpace(input)
	// Convierto jel texto a número y lo almaceno en la variable [ano]
	var err error
	ano, err = strconv.Atoi(input)
	if err != nil {
		// Si la conversión falla, informo al usuario del error.
		fmt.Println("No es una matrícula valída", err)
	}

	// Llamo al constructor del vehículo.
	v := vehiculo.NuevoVehiculo(tipo, marca, modelo, matricula, ano)

	// Muestro la iformación del vehículo.
	v.ImprimirInfo()

	// Utilizo un menú para ejecutar la acción concreta del vehículo.
	for !salir {
		// Imprimo el menú.
		fmt.Print(`
1. Arrancar el vehículo.
2. Frenar el vehículo.
3. Tocar el claxon.
4. Salir.
Seleccione una opción: `)
		// Leo la opción seleccionada.
		fmt.Scanln(&opcion)

		// Ejecuto la opción seleccionada por el usuario.
		switch opcion {
		case 1:
			v.Arrancar() // Llamada al método arrancar del objeto vehículo
		case 2:
			v.Frenar() // Llamada al método frenar del objeto vehículo
		case 3:
			v.Claxon() // Llamada al método claxon del objeto vehículo
		case 4:
			// Salgo del bucle
			salir = true
		default:
				fmt.Println("Opción, como tu opinión, inválida")
		}
	}
}
