// Indico el paquete en el que se encuentra el archivo.
package ejercicios

/*
###########
#ENUNCIADO#
###########

1.	Crear una aplicación en go, que imprima la mítica frase “hola, Mundo! En la consola.

2.	Imprimir en otra línea tu nombre y en otra distinta tus apellidos.

3.	Generar el archivo “main.exe”

*/

// Importo el paquete para tratar con la consola.
import "fmt"

// Declaro y defino una función para imprimir el nombre.
func MostrarNombre() {
	// Imprimo la frase requqerida.
	fmt.Println("Hola Mundo!")
	// Imprimo el nombre.
	fmt.Println("Armando")
	// Imprimo los apellidos.
	fmt.Println("Casitas Rojas")
}
