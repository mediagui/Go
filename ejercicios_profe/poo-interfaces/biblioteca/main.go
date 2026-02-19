package main

// Importaciones.
import (
	// Importo la librería para crear los objetos.
	"biblioteca/libro"
)

/*
En Go, la programación orientada a objetos (POO), es un enfoque que utiliza estructuras y métodos para realizar una representación de un objeto del mundo real en nuestro código. Aunque Go no se considera totalmente orientado a objetos, ofrece varias características que nos pueden ayudar a simularla (estructuras, métodos, interfaces y composición)

# Estructuras:

	Las estructuras en Go son similares a las clases en otros lenguajes de POO. Las estructuras se utilizan para representar objetos y tienen campos para almacenar los datos.

# Métodos:

	Los métodos en Go son funciones que operan en una estructura en particular. Se utilizan para añadir comportamientos a las estructuras y para definir operaciones específicas para un objeto.

# Encapsulamiento:

	Aunque Go no tiene modificadores de acceso públicos o privados, se puede lograr el encapsulamiento utilizando conveciones de nomenclatura de variables y estructuras.

# Composición (simula herencia):

	Go utiliza la composición para modelar relaciones entre objetos. Esto implica la creación de una estructura que contiene otras estructuras.

# Interfaces:

	Las interfaces en Go son tipos abstractos que definen un conjunto de métodos. Permiten que diferentes tipos implementen el mismo conjunto de métodos, lo que permite polimorfismo y abstracción
*/

// Función principal.
func main() {
	miLibro := libro.NuevoLibro("El Señor de las Moscas", "William Golding", 288)

	// var miLibro = libro.Libro{
	// 	Titulo:  "El Señor de las Moscas",
	// 	Autor:   "William Golding",
	// 	Paginas: 288,
	// }

	miLibro.ImprimirInfo()
}
