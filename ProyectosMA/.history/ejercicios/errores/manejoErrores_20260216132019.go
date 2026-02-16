// Se pide desarrollar una aplicación que permita a los usuarios guardar y administrar contactos en un archivo.
// Hay que recordar manejar los posibles errores que puedan ocurrir.
package errores

import "os"

func main() {}

func createNotepad() {
	fichero, error := os.OpenFile("agenda.txt")
	checkError(error)
	defer func() {
		checkError(fichero.Close())
	}()
}

func checkError(e error) {
	if e != nil {

	}
}
