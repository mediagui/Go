// Se pide desarrollar una aplicación que permita a los usuarios guardar y administrar contactos en un archivo.
// Hay que recordar manejar los posibles errores que puedan ocurrir.
package errores

import "os"

func main() {
	createNotepad()

}

func createNotepad() *os.File {
	fichero, error := os.OpenFile("agenda.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	checkError(error)

	defer func() {
		checkError(fichero.Close())
	}()
	panicRecover()

	return fichero

}

func checkError(e error) {

	defer func() {
		if e != nil {
			panic("Entro en pánico. Error grave.")
		}
	}()

}

func panicRecover() {
	defer func() {
		if recovering := recover(); recovering != nil {
			println("Recuperamos la calma después de ", recovering)
		}
	}()
}
