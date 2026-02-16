// Se pide desarrollar una aplicación que permita a los usuarios guardar y administrar contactos en un archivo.
// Hay que recordar manejar los posibles errores que puedan ocurrir.
package errores

import (
	"fmt"
	"log"
	"os"
)

type contactStruct struct {
	name    string
	surname string
}

func main() {

	contacts := map[int]contactStruct{}
	fillContactMap(contacts)

	agenda := createAgenda()
	saveContactsToAgenda(contacts, agenda)

}

func saveContactsToAgenda(contacts map[int]contactStruct, agenda *os.File) {
	for id, contact := range contacts {
		line := fmt.Sprintf("%d: %s %s\n", id, contact.name, contact.surname)
		n, err := agenda.WriteString(line)
		log.Printf("n %d bytes written\n", n)

		checkError(err)

	}
}

func createAgenda() *os.File {

	fichero, error := os.OpenFile("agenda.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	log.Printf("File %s created\n", fichero.Name())

	checkError(error)

	func() {
		checkError(fichero.Close())
	}()

	panicRecover()

	return fichero

}

func fillContactMap(contacts map[int]contactStruct) {
	for i := 0; i < 10; i++ {
		contacts[i] = buildContact(i)
	}
}

func buildContact(id int) contactStruct {

	return contactStruct{
		name:    fmt.Sprintf("Nombre %d", id),
		surname: fmt.Sprintf("Apellido %d", id),
	}

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
