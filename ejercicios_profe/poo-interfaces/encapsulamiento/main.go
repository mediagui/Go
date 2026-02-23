package main

// Importaciones.
import (
	// Importo la librería para crear los objetos.
	"biblioteca/libro"
	"fmt"
)

/*
En Go, la programación orientada a objetos (POO), es un enfoque que utiliza estructuras y métodos para realizar una representación de un objeto del mundo real en nuestro código. Aunque Go no se considera totalmente orientado a objetos, ofrece varias características que nos pueden ayudar a simularla (estructuras, métodos, interfaces y composición)

# Encapsulamiento:

	Aunque Go no tiene modificadores de acceso públicos o privados, se puede lograr el encapsulamiento utilizando conveciones de nomenclatura de variables y estructuras.

	Es el principio por el cual una clase u objeto oculta sus datos internos y solo permite acceder a ellos mediante métodos controlados.

	En otras palabras: "Encapsulasr es guardar algo dentro de una caja y decidir qué partes se pueden tocar por los demás"

	¿Por qué es útil?

	1. Protege los datos.

		Evita que otras partes del programa amodifiquen valores internos de forma accidental o incorrecta.

	2. Controla cómo se accede a los datos.

		Puedes decidir si permites lectura, escritura o ambas sobre los atributos de una estructura.

	3. Reduce errores y efectos secundarios.

		Si los datos están protegidos, el resto del programa no puede romperlos sin querer.

		Esto hace que el código sea:

			Más estable.
			Más fácil de depurar.
			Más predecible.

	4. Permite cambiar la implementación sin romper el código.

		En un ejemplo en el que tengamos un dato edad de tipo [int], que queramos en un futuro calcular a partir de una fecha de nacimiento, si los datos están encapsulados, podemos cambiar la lógica interna sin que el resto del programa se vea afectado.
*/

// Función principal.
func main() {
	miLibro := libro.NuevoLibro("El Señor de las Moscas", "William Golding", 288)

	/*
		Al estipular los elementos como privados, no puedo acceder a modificarlos de esta manera.
	*/
	// libro := libro.Libro{"El Señor de las Moscas","William Golding",288}

	// var miLibro = libro.Libro{
	// 	Titulo:  "El Señor de las Moscas",
	// 	Autor:   "William Golding",
	// 	Paginas: 288,
	// }

	miLibro.ImprimirInfo()

	fmt.Println(`
#####################
#CON ENCAPSULAMIENTO#
#####################
	`)

	// Creo un nuevo libro para ejemplificar el concepto de encapsulamiento.
	miLibroDos := libro.NuevoLibro("El Diario de Ana Frank", "Ana Frank", 1945)

	miLibroDos.ImprimirInfo()

	miLibroDos.SetTitulo("El Diario de Patricia")

	fmt.Println(miLibroDos.GetTitulo())
	fmt.Println(miLibroDos.GetAutor())
	fmt.Println(miLibroDos.GetPaginas())

}
