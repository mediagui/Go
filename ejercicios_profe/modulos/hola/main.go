// Indico que pertenece al paquete principal
package main

// Importo la librería [fmt]
import (
	"fmt"
	"log"

	"github.com/nox0V0saw/agradecimientos"
)

// Función principal.
func main() {
	// Establezco un prefijo para los logs.
	log.SetPrefix("[ERROR en agradecimientos] ")
	// Deshabilito la impresión de fecha y hora por defecto.
	log.SetFlags(0)

	// Creo un slice con varios nombres.
	nombres := []string{"Eustaquio", "Evaristo", "Candida", "Ramona", "Calpurnia", "Modesta"}

	// Imprimo un saludo en la consola.
	mensajes, err := agradecimientos.HolaVarias(nombres)

	// Compruebo el error y lo manejo.
	if err != nil {
		log.Fatal(err)
	}
	for mensaje, _ := range mensajes {
		fmt.Println(mensaje)
	}

}
