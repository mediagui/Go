// Definición del paquete principal.
package main

// Importaciones
import (
	"bufio"                             // Paquete para la lectura eficiente de datos de entrada.
	"concesionarioComposicion/vehiculo" // Importo el paquete local de gestión de vehículos.
	"fmt"                               // Paquete estándar para entrada y salida formateada.
	"os"                                // Paquete para interactuar con el sistema operativo (entrada estándar)
	"strconv"                           // Paquete para converión de cadenas a otros tipos (string a int)
	"strings"                           // Paquete para manipular cadenas de texto.
)

/*
###########
#ENUNCIADO#
###########

Modificar el programa que permite gestionar un concesionario, creando automóviles, de manera que exista una estructura “padre” general para todos los vehículos y cada “hijo”, tenga sus atributos específicos:

v Coche:

o  Número de puertas.

o  Motor.

o  ¿Descapotable?

v Moto

o  Cilindrada

o  Potencia

v Camión

o  MMA (Masa Máxima Autorizada)

o  Tipo de carga (cisterna, animales, productos alimenticios, etc…)
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
4. Modificar matrícula.
5. Modificar año.
6. Mostrar la información del vehículo.
7. Salir.
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
		case 4: // Modificar la matrícula con set
			fmt.Println("Ingrese la nueva matrícula")
			nuevaMatricula, _ := rd.ReadString('\n')
			nuevaMatricula = strings.TrimSpace(nuevaMatricula)

			// Manejo el posible error.
			if err := v.SetMatricula(nuevaMatricula); err != nil {
				fmt.Println("[ERROR]: ", err)
			} else { // En caso de que no haya habido un error, muestro un mensaje de confirmación.
				fmt.Println("Matrículo actualizada correctamente.")
			}
		case 5: // Modificar el año con set
			fmt.Println("Ingrese el año modificado:")
			inputAno, _ := rd.ReadString('\n')
			inputAno = strings.TrimSpace(inputAno)
			nuevoAno, err := strconv.Atoi(inputAno)

			// Manejo el posible error.
			if err != nil {
				fmt.Println("[ERROR]: Ingrese un número válido")
				continue
			}

			// Manejo el posible error al introducir el nuevo año.
			if err := v.SetAno(nuevoAno); err != nil {
				fmt.Println("[Error] de validación: ", err)
			} else {
				fmt.Println("Año actualizado correctamente.")
			}
		case 6: // Muestro la información de vehículo.
			v.ImprimirInfo()
		case 7:
			// Salgo del bucle
			salir = true
		default:
			fmt.Println("Opción, como tu opinión, inválida")
		}
	}
}
