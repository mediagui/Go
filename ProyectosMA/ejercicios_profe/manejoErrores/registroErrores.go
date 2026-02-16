// Paquete del proyecto.
package manejoErrores

// Importación de paquetes
import (
	"fmt"
	"log"
	"os"
)

func RegistarErrores() {
	/*
		El paquete [log] se utiliza para el registro básico de mensajes en la aplicación. Proporciona funconalidades simples para imprimir mensajes de registro en la salida estándar y en la salida estándar de errores.

		[log.Print]
			-> Imprime el error en la consola.
		[log.Fatal]
			-> Imprime el error en la consola y termina la ejecución.
		[log.SetPrefix]
			-> Se puede utilizar para añadir un prefijo a los mensajes de registro del programa.

	*/

	// Indico el prefijo que va a tener cada log.
	log.SetPrefix("RegistarErrores(): ")

	log.Print("Este es un mensaje de registro")
	log.Println("Este es otro mensaje de registro.")
	// log.Fatal("¡OYE, soy un registro de errores")
	fmt.Print("¿Puedes verme?")

	/*
		Abro un archivo, si no existe.

		["info.log"]
			-> Nomrbe del archivo.
		[os.O_CREATE]
			-> Crea el archivo si no existe.
		[os.O_APPEND]
			-> Escribe al final del archivo.
		[os.O_WRONLY]
			-> Abre el archivo solo para escritura.
		[0644]
			-> Permisos para el archivo (0644) rw-r--r--
	*/
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// Manejamos el error, en caso de que ocurra uno a la hora de abrir/crear el archivo.
	if err != nil {
		// En caso de que se produzca el error, interrumpo la ejecución y devuelvo el mensaje.
		// fmt.Println("asdfasdfas")
		log.Fatal(err)
	}

	// [defer] garantiza que se cierre el archivo al final del main, incluso aunque ocurra un error más adelante.
	defer file.Close()

	// Configuro el logger para que escriba en el archivo, en vez de en la consola.
	log.SetOutput(file)

	// Escribo un mensaje en el archivo de log.
	log.Print("¡Oye! Mi cuerpo pide salasa")
}
