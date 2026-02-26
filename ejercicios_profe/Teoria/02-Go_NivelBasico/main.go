/*
Para poder ejecutar correctamente el proyecto, nos moveremos con el cmd a la carpeta donde se encuentra nuestro proyecto y ejecutaremos el comando [go mod init nombreCarpetaTrabajo].

Para ejecutar el programa, utilizaremos el comando [go run nombrePrograma.go]
*/
// Indico el paquete al que pertenece el programa. Si no se encuentra dentro de ninguna carpeta, se suele dejar por defecto, en este caso, el nombre main.
package main

// Importo el paquete ["fmt"] (format).
// Contiene las funciones necesarias para imprimir text y leer entradas por consola.
import (
	// "fmt"
	// "ProyectosGo/ejercicios"
	"ProyectosGo/variables"
)

// La función [main()] es el punto de entrada obligatorio de la aplicación.
// No recibe parámetros ni devuelve valores.
func main() {
	// Imprimo la mítica frase "hola, Mundo!" en la consola.
	// Esto lo haremos con la función [Println()] del paquete [fmt]. Esta función nos permite imprimir mensajes en la consola.
	// fmt.Print("hola Mundo! 🐱‍👤")
	// Llamo a la función que imprime lo requerido en el ejercicio 1.
	//ejercicios.MostrarNombre()
	// Llamo a la función para imprimir los tipos de datos.
	variables.MostrarDatos()
}

// Si ejecutamos el comando [go build main.go], se creará un [main.exe] que nos permitirá ejecutar el programa en cualquier parte.
