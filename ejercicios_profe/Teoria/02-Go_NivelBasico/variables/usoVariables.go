// Defino el paquete en el que se encuentra el archivo.
package variables

// Importación de paquetes.
import (
	"fmt"
)

// Función en la que trataré con distintos tipos de variables y datos e imprimiré sus valores.
func TratoVariables() {
	/*
		Las variables son un tipo de dato de memoria temporal, en go las variables son tipadas, por lo que debemos definir un tipo de dato de variables.

		Una manera que tenemos de declarar una variable es con la sintaxis:

		var nombreVariable tipoDato
	*/
	// Declaro la variable [nombreUno] y le indico el típo de dato.
	var nombreUno string

	// Declaro dos variables al mismo tiempo, que comparte tipo de dato.
	var nombreDos, apellidoUno string
	// Declaro una variable para almacenar la edad.
	var edad int

	// Defino la variable [nombreUno], asignándole un valor de tipo [string].
	nombreUno = "Víctor"
	// Defino las variables [nombreDos] y [apellidoUno], asignándoles un valor de tipo [string].
	nombreDos = "Luis"
	apellidoUno = "Glaría"
	// Asigno un valor a la variable [edad]
	edad = 29

	// Declaro 3 variables de tipos distintos.
	var (
		pais, ciudad string
		cPostal      int
	)

	// Declaro y defino 3 variables al mismo tiempo.
	var (
		dni         string = "15988765a"
		tfno        int    = 689452566
		estadoCivil string = "Soltero"
	)

	// Asigno valores a las variables [pais], [ciudad] y [cPostal]
	pais = "España"
	ciudad = "Madriz"
	cPostal = 28932

	// Imprimo el valor de la variable [nombreUno]
	fmt.Println(nombreUno)
	// Imprimo el valor de las variables [nombreDos] y [apellidoUno]
	fmt.Println(nombreDos)
	fmt.Println(apellidoUno)
	// Imprimo el valor de la variable [edad]
	fmt.Println(edad)

	// Imprimo los valores de mi dirección.
	fmt.Println("Pais:", pais, "Ciudad:", ciudad, "Código Postal:", cPostal)

	// Imprimo más valores sobre mi.
	fmt.Println("DNI:", dni, "Nº Tfno:", tfno, "Estado Civil:", estadoCivil)
}
