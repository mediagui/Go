// Paquete del proyecto.
package manejoErrores

// Importación de paquetes
import (
	"fmt"
	"os" // Paquete que trabaja con archivos y el sistema operativo
)

func EjpDefer() {
	/*
		Creo un archivo [hola.txt]. En este punto [os] deuvelve dos valores:

		1 - Un puntero al archivo creado.
		2 - Un error si algo salió mal.
	*/
	file, err := os.Create("hola.txt")
	// Aplazo hasta el final de la ejecución el cierre del archivo.
	defer file.Close()
	// Compruebo si ha ocurrido algún error al crear el archivo.
	if err != nil {
		fmt.Println(err)
		return
	}

	// Escribo en el archivo creado. [file.Write([]byte("Texto a escribir"))] recibe un slice de [bytes].
	_, err = file.Write([]byte("Hola caracola"))
	if err != nil {
		// En caso de que haya un error al escribir, muestro el mensaje correspondiente y termino la ejecución del programa.
		fmt.Println(err)
		return
	}
}
