// Indico el paquete al que pertenece el archivo.
package paquetefmt

import "fmt"

/*
###########
#VARIABLES#
###########
*/
var (
	nombre        string = "Víctor"
	edad          int    = 29
	apellidos     string
	apellidoUno   string
	apellidoDos   string
	anoNacimiento int
)

/*
	El paquete [fmt], no solo permite mostrar texto en consola, sino que también nos permite introducir datos en la misma.

	Si queremos saber más sobre este paquete, podemos visitar la documentación oficial en el siguiente enlace:

	https://pkg.go.dev/fmt
*/

// Defino la función con la que trataremos con las distintas opciones que nos da el paquete [fmt]. Esta función no recive ni devuelve ningún valor.
func FuncionFmt() {
	// Como vemos, si utilizo [Println], se imprime el texto y se añade un salto de línea al final.
	fmt.Println("Mensaje con [Println]")
	// Sin embargo, si utilizo [Print], no se añade salto de línea al final.
	fmt.Print("Mensaje con [Print] normal")
	fmt.Print("Mensaje con [Print] continuado\n") // Como vemos, se muestran los dos mensajes seguidos. Para que el siguiente mensaje se mostrase en una línea separada, añadimos un salto de línea con [\n]

	// Como hemos visto hasta ahora, con [Printf], podemos formatear el texto, juntando variables con distintos tipos de datos en nuestros [strings].

	// Muestro un texto con los valores de las variables [nombre] y [edad]
	fmt.Printf("Mi nombre es [%s] y tengo [%d] años.\n", nombre, edad)

	/*
		Otra forma de hacerlo es utilizando [Sprintf], que en lugar de imprimir el trxto en consola, lo devuelve como una cadena de texto y nos permite almacenarla en una variable.
	*/

	// Declaro una variable y la defino con un valor devuelto po [Sprinf].
	mensajeBienvenida := fmt.Sprintf("Bienvenido [%s], de [%d] años.", nombre, edad)

	// Muestro el mensaje almacenado en la variable. Al imprimir con [Println], se añade un salto de línea al final.
	fmt.Println(mensajeBienvenida)

	/*
		El paquete [fmt] también nos permite introducir datos por consola utilizanso la función [Scanln]. Normalmente, se suele mostrar un mensaje por consola indicando al usuario que introduzca los datos, y a continuación se utiliza [Scanln] para capturar la entrada del usuario. También se suele dejar el cursor en la misma línea que el mensaje, para mejor experiencia del usuario.
	*/
	// Solicito al usuari que introduzca sus apellidos.
	fmt.Print("Por favor, introduce tu primer apellido: ")
	fmt.Scanln(&apellidos) // Capturo la entrada del usuario y la almaceno en la variable [apellidos]. Es importante pasar la dirección de memoria de la variable utilizando el operador [&].

	// Solicito al usuario que introduzca su año de nacimiento.
	fmt.Print("Por favor, introduzca su año de nacimiento: ")
	// Capturo la entrada del usuario y la almaceno en una variable.
	fmt.Scanln(&anoNacimiento)

	// Muestro un mensaje con un saludo completo al usuario.
	fmt.Printf("Hola [%s] [%s], naciste en el año [%d], por lo que tienes [%d] años.\n", nombre, apellidos, anoNacimiento, edad)

	// Podría capturar dos valores a la vez, indicando dos variables en la función [Scanln]
	fmt.Print("Por favor, introduce tus dos apellidos, separados por un espacio: ")
	// Capturo la entrada del usuario y la almaceno en las variables [apellidoUno] y [apellidoDos].
	fmt.Scanln(&apellidoUno, &apellidoDos)
	// Muestro un mensaje con un saludo completo al usuario.
	fmt.Printf("Hola [%s] [%s] [%s], naciste en el año [%d], por lo que tienes [%d] años.", nombre, apellidoUno, apellidoDos, anoNacimiento, edad)

	/*
		NOTA: Si se introduce un valor que no coincide con el tipo esperado, se producirá un error en tiempo de ejecución. Para imprimir cualquier tipo de dato con el método [format], se puede utilizar el especificador %v.
	*/
	// Imprimo una frase formateada, sin especificar el tipo de dato a introducir.
	fmt.Printf("Imprimiendo cualquier tipo de dato con %%v: %v tiene %v años.\n", nombre, edad) // Pero esto se  recomienda cuando no se conoce el tipo de dato que se va a imprimir.

	fmt.Println("################################")
	fmt.Println("#IMPRESIÓN DE VARIABLES Y TIPOS#")
	fmt.Println("################################")
	/*
		Con el modificador [%T], podemos conocer el tpo de dato de una variable.
	*/
	// Imprimo el tipo de dato de todas las variables utilizadas junto con su valor.
	fmt.Printf("Variable [nombre]\nTipo dato: [%T]\nValor[%v]", nombre, nombre)
	fmt.Printf("Variable [edad]\nTipo dato: [%T]\nValor[%v]\n", edad, edad)
	fmt.Printf("Variable [apellidos]\nTipo dato: [%T]\nValor[%v]\n", apellidos, apellidos)
	fmt.Printf("Variable [apellidoUno]\nTipo dato: [%T]\nValor[%v]\n", apellidoUno, apellidoUno)
	fmt.Printf("Variable [apellidoDos]\nTipo dato: [%T]\nValor[%v]\n", apellidoDos, apellidoDos)
	fmt.Printf("Variable [anoNacimiento]\nTipo dato: [%T]\nValor[%v]\n", anoNacimiento, anoNacimiento)
}
