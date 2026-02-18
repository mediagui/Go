// Indico el paquete al que pertenece el programa.
package estructuras

// Importaciones
import "fmt"

/*
#########
#STRUCTS#
#########

Es una estructura de datos que nos permite definir un tipo de dato personalizado, compuesto por varios campos con nombres y tipos epecíficos.

#SINTAXIS#
type NombreStruct struct {
	Campo 1 Tipo
	Campo 2 Tipo
	Campo n Tipo
}

#Deseglose#

→ Type:
	Es la palabra reservada que indica que estamos definiendo un tipo.
→ NombreStruct:
	Es el nombre que le damos a la estructura.
→ struct:
	Es la palabra reservada que indica que estamos definiendo una estructura.
→ Campo1, Campo2, CampoN:
	Son los nombres de los campos de la estructura.
→ Tipo:
	Es el tipo de dato que puede tener cada campo.
*/

// Creo una estructura llamada "Persona". Esta tendrá como atributos: nombre, edad y email.
type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

// Zona de declaración de variables.
var (
	// Variables para almacenar los códigos de colores.
	codigoRojo   = "\033[38;2;255;0;0m"
	codigoVerde  = "\033[38;2;0;255;0m"
	codigoAzul   = "\033[38;2;0;0;255m"
	codigoBlanco = "\033[38;2;255;255;255m"
	fondoBlanco  = "\033[48;2;255;255;255m"
	fondoNegro   = "\033[48;2;0;0;0m"

	// Variable para almacenar el reset de los colores.
	reset = "\033[0m"
)

// Función para demostrar el uso de estructuras.
func ExplicacionEstructuras() {
	// fmt.Print("FUNCIONO JODER")
	// Una vez creada la estructura, podemos crear un dato (crear una instancia o instanciar) de tipo Persona.
	var persona1 Persona

	// Si intento imprimir la [persona1], veré que solo se muestran los datos por defecto, de los atributos que hemos declarado en la estructura.
	fmt.Println(persona1)

	/*
		Si quiero imprimir los atributos de la [persona1], debo acceder a cada uno de ellos. Si trato de imprimirlos sin asignarles ningún valor, veremos que devuelve un valor por defecto.
	*/
	fmt.Println(persona1.Nombre)
	fmt.Println(persona1.Edad)
	fmt.Println(persona1.Email)

	// Una manera de definir los atributos de un dato, es llamando a cada atributo 1 por 1, asignándoles un valor.
	persona1.Nombre = "Juan"
	persona1.Edad = 30
	persona1.Email = "juan@juan.com"

	// Imprimo los datos de [persona1]
	fmt.Println(fondoNegro+codigoVerde+"Persona 1: ", persona1, reset)

	/*
		Otra manera de definir los atributos de un dato, es llamando a la estructura y asignándole un valor para cada atributo.
	*/
	persona2 := Persona{Nombre: "Ana", Edad: 25, Email: "ana@gmail.com"}
	// Imprimo los datos de [persona2]
	fmt.Println(fondoNegro+codigoBlanco+"Persona 2: ", persona2, reset)

	/*
		Otra manera de definir los atributos de un dato, es llamando a la estructura y asignándole un valor, sin indicar el nombre del atributo, pero esto hay que hacerlo en el orden en que se definen.
	*/
	persona3 := Persona{"Pepe", 28, "pepe@gmail.com"}
	// Imprimo los datos de [persona3]
	fmt.Println(fondoNegro+codigoRojo+"Persona 3: ", persona3, reset)

	// También puedo modificar los atributos de un dato, llamando a la estructura y asignándole un valor para cada atributo.
	persona3.Nombre = "José Jesús"

	// Muestro únicamente el nombre y la edad de la [persona3].
	fmt.Println(fondoNegro+codigoVerde+"Persona 3: ", persona3.Nombre, persona3.Edad, reset)
}
