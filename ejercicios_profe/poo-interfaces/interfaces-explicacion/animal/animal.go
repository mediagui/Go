package animal

/*
#IMPORTACIONES#
*/
import "fmt"

/*
############
#INTERFACES#
############

En go, las interfaces son un mecanismo poderoso para definir qué métodos deben ser implementados por otros tipos de datos (structs).

#Sintaxis#

type NombreInterfaz interface {
	// Conjunto de métodos que deben ser implementados por cualquier tipo que quiera cumplir con esta interface (contrato)
	Metodo1()
	Metodo2()
	MetodoN()
}
*/

/*
1. El Contrato (La interface)

Aquí defino el estándar. Indicamos que, cualquier cosa que quiera ser un [Animal] en este programa, debe tener un método [Sonido()]
*/
type Animal interface {
	Sonido()
}

/*
2. Las Implementaciones (Los Actores)
*/

// Defino la estructura [Perro]. Solo tiene datos (Nombre)
type Perro struct {
	Nombre string
}

// Hago que el perro "firme" el contrato. Al crear un método [Sonido()], el Perro se conviete automáticamente en un Animal
func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace guau guau")
}

// Defino la estructura [Gato]. Solo tiene datos (Nombre)
type Gato struct {
	Nombre string
}

// Hago que el perro "firme" el contrato. Al crear un método [Sonido()], el Perro se conviete automáticamente en un Animal
func (p *Gato) Sonido() {
	fmt.Println(p.Nombre + " hace miau")
}

/*
3. La función Gnérica (El usuario de la interface)

A esta función no le importa si le paso un perro o un gato, solo exige que el objeto cumpla el contrato [Animal]

Esto es lo que permite que el código sea más flexible y no se pueda romper si añadimos nuevos animales.
*/
func HacerSonido(animal Animal) {
	animal.Sonido() // Go sabe exáctamente a qué [Sonido()] llamar en tiempo de ejecución
}
