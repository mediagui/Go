// Paquete al que pertenece el programa.
package punterosmetodos

// Importaciones
import "fmt"

// Creo una estructura de una persona.
type Persona struct {
	Nombre            string
	Edad              int
	CorreoElectronico string
}

// Método para que la persona salude.
func (p *Persona) DiHola() {
	fmt.Println("Hola, mi nombre es: ", p.Nombre)
}

// Función para la explicación de punteros y métodos
func ExplicacionPunterosMetodos() {
	fmt.Println(`
####################
#PUNTEROS Y MÉTODOS#
####################
`)

	/* --- PARTE 1: ASPECTOS BÁSICOS --- */

	// Declaro y defino una variable normal.
	var numero int = 10

	// Creo una variable que almacenará el puntero a un dato de tipo [int]. El [*], indica que la variable no va a guardar un número, sino la dirección en memoria del número.
	var puntero *int

	// Para recuperar esta dirección en memoria, utilizamos [&]. Esto es como pedirle al sistema que nos devuelva la dirección en la que "vive" el valor de la variable.
	puntero = &numero

	fmt.Println("--- 1. Ejemplo Básico PUNTERO ---")
	// Muestro el valor de la variable [numero]
	fmt.Println("El valor de la variable [numero] es: ", numero)

	// Muestro el valor de [puntero]. Al hacerlo, veremos la dirección en memoria, un código hexadecimal como 0x000....
	fmt.Println("La dirección en memoria es: ", puntero)

	// Para ver qué hay dentro de esta direción, utilizamos el asterisco [*], delante del puntero. Esto se llama "desreferenciar".
	fmt.Println("Si seguimos la dirección del puntero, encontramos el valor: ", *puntero)

	/*
		--- Parte 2: Error común (nil pointer) ---
	*/

	fmt.Println("\n--- 2. ¡Cuidado con los errores! ---")

	// Un error muy típico es crear un puntero, pero no indicarle a dónde debe apuntar.
	// En Go, si no le asignamos una dirección, el puntero vale [nil] (está vacío).
	var punteroVacio *int

	fmt.Println("El valor actual del puntero [punteroVacio], es: ", punteroVacio)

	/*
		¡ATENCIÓN! Si intentamos hace [fmt.Println(*punteroVacio)], el programa EXPLOTARÁ (daría un [panic]) por que le estaríamos pidiendo que vaya a una dirección que no existe.

		fmt.Println(*punteroVacio) // <- El programa explota con esto.
	*/

	// Para evitar esto, se suele utilizar de la siguiente manera.
	if punteroVacio == nil {
		fmt.Println("El puntero está vacío, no podemos entrar a ver su interior todavía.")
	}

	/*
		#########
		#MÉTODOS#
		#########

		Se pueden definir los métodos en estructuras para realizar operaciones en las variables de la estructur. Los métodos se definen como funciones que tienen un receptor, que es un puntero o una variable de la estructura.

		#Sintaxis#

		func (nombreObjeto *Puntero) NombreMetodo(){
			// Bloque de instrucciones
		}
	*/

	// Defino una persona.
	persona1 := Persona{"Diego", 68, "tuDieguito69@hotmail.es"}
	fmt.Println(persona1)

	persona1.DiHola()

}
