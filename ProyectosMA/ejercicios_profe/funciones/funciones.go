// Indico el paquete al que pertenece el programa.
package funciones

// Importaciones.
import "fmt"

// Zona de declaración de variables.
var (
	// Variables para almacenar los códigos de colores.
	codigoRojo   = "\033[38;2;255;0;0m"
	codigoVerde  = "\033[38;2;0;255;0m"
	codigoAzul   = "\033[38;2;0;0;255m"
	codigoBlanco = "\033[38;2;255;255;255m"
	fondoBlanco  = "\033[48;2;255;255;255m"
	fondoNegro   = "\033[48;2;0;0;0m]"

	// Variable para almacenar el reset de los colores.
	reset = "\033[0m"
)

func ExplicacionFunciones() {
	/*
		############
		#FUNCIONES #
		############

		Las funciones son útiles para ordenar nuestro código y también para reutilizar un bloque de instrucciones multiples veces.codigoAzul

		#SINTAXIS#

		func nombreFunción(parametros) (retorno) {
			// Cuerpo de la función (bloque de instrucciones)
		}
	*/
	// Llamo a la función [saludar()]
	saludar()

	// Pido al usuario su apellido y lo almaceno en una variable.
	fmt.Print("Ingrese su apellido: ")
	var apellido string
	fmt.Scan(&apellido)
	// Llamo a la función [retornarApellidos()] y le paso como parámetro el apellido introducido por el usuario.
	retornarApellidos(apellido)
	// Almaceno en una variable la edad del usuario, mediante la función, indicando una edad inicial.
	// Llamo a la función que devuelve la edad
	edad := retornarEdad(9)
	fmt.Println("Su edad es: ", edad)

	// Llamo a la función [calc()]
	sum, mul := calc(2, 3)
	fmt.Printf(`
La suma es: %d
La multiplicación es: %d
	`, sum, mul)
}

/*
Creo la función [saludar]. Esta función solicitará al usuario el nombre y lo imprimirá. No recibe ni devuelve ningún valor.
*/
func saludar() {
	fmt.Println("Hola desde la función [saludar()]")

	// Declaro la variable [nombre] y la defino con un nombre que ingrese el usuario.
	var nombre string
	fmt.Print("Ingrese su nombre\n-> ")
	fmt.Scan(&nombre)
	fmt.Printf("%s %s Hola %s, te saludo desde la función saludar%s\n", fondoNegro, codigoVerde, nombre, reset)
}

// Declaro la función [retornarApellidos]. Esta espera como parámetro un string [fapellido] e imprimie el apellido. No devuelve ningún valor.
func retornarApellidos(fapellido string) {
	fmt.Printf("Su apellido es: %s \n", fapellido)
}

// Defino la función [retornarEdad()], este espera como parámetro un dato de tipo [int] y devolverá un dato de tipo [int]
func retornarEdad(fedad int) int {
	// Hago una función recursiva, que se ejecutará hasta que el usuario sea mayor de edad.
	if fedad < 18 {
		fmt.Println("Usted es menor de edad, vuelva a intentarlo.")
		fmt.Print("Ingrese su edad nuevamente: ")
		fmt.Scan(&fedad)
		// Tengo que devolver el valor de la función recursiva, por que sino, no se actualiza el valor del parámetro [fedad], en caso de que sólo haga una llamada a la función.
		return retornarEdad(fedad)
	} else {
		fmt.Println("Usted es mayor de edad")
		return fedad
	}
}

/*
También puedo indicar exactamente qué variables se van a retornar.
*/
// Función que recibe dos operandos, realiza la suma y la multiplicación con dichos operandos y las devuelve.

// func calc(fop1, fop2 int) (int, int) {
// 	fsum := fop1 + fop2
// 	fmul := fop1 * fop2
// 	return fsum, fmul
// }

func calc(fop1, fop2 int) (fsum int, fmul int) {
	// fsum := fop1 + fop2
	// fmul := fop1 * fop2
	// Si lo hacemos de esta manera, no es necesario declarar las variables [fsum] y [fmul], sino que simplemente les asignamos los valores.
	fsum = fop1 + fop2
	fmul = fop1 * fop2
	return // fsum, fmul
}
