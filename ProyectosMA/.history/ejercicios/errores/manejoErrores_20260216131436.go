// Se pide desarrollar una aplicación que permita a los usuarios guardar y administrar contactos en un archivo.
// Hay que recordar manejar los posibles errores que puedan ocurrir.
package errores

import "os"

func main() {}

func createNotepad(){
	fichero, error := os.Create("agenda.txt")
	checkError(error)
	closeError := func(){ defer return fichero.Close()}()
	checkError(closeError)


}

func checkError(e error){}
