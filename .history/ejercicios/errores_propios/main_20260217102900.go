// Se pide desarrollar una aplicación que permita a los usuarios guardar y administrar contactos en un archivo.
// Hay que recordar manejar los posibles errores que puedan ocurrir.
package main

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

	log.Println("Starting...")

	contacts := map[int]contactStruct{}
	fillContactMap(contacts)

	agenda := createAgenda()
	saveContactsToAgenda(contacts, agenda)

	// =======

	agenda = openAgenda(true)

	showContentOf(agenda)

	panicRecover()

	os.Exit(0)

}

func showContentOf(agenda *os.File) {

	agendaContent, err := os.ReadFile(agenda.Name())
	checkError(err)
	fmt.Printf("Contenido: \n\t%v\n", agendaContent)

}

func saveContactsToAgenda(contacts map[int]contactStruct, agenda *os.File) {

	agenda, err := os.OpenFile("agenda.txt", os.O_APPEND|os.O_WRONLY, 0644)
	checkError(err)
	defer checkError(agenda.Close())

	for id, contact := range contacts {
		line := fmt.Sprintf("%d: %s %s\n", id, contact.name, contact.surname)
		n, err := agenda.WriteString(line)

		log.Printf("%d bytes written to %s file\n", n, agenda.Name())

		checkError(err)

	}
}

func createAgenda() *os.File {

	log.Println("Creating file...")

	fichero, error := os.Create("agenda.txt")

	func() {
		checkError(fichero.Close())
		log.Printf("File %s properly closed.\n", fichero.Name())
	}()

	log.Printf("File %s created.\n", fichero.Name())

	checkError(error)

	panicRecover()

	return fichero

}

func openAgenda(asReadOnly bool) *os.File {
	var perm int
	if asReadOnly {
		perm = os.O_RDONLY
	} else {
		perm = os.O_WRONLY
	}

	agenda, err := os.OpenFile("agenda.txt", perm, 0644)
	checkError(err)

	return agenda

}

func fillContactMap(contacts map[int]contactStruct) {
	log.Println("Filling contact map...")
	for i := range 10 {
		contacts[i] = buildContact(i)
		log.Printf("Added contact %d\n", i)
	}
}

func buildContact(id int) contactStruct {

	log.Printf("Building contact %d\n", id)

	return contactStruct{
		name:    fmt.Sprintf("Nombre %d", id),
		surname: fmt.Sprintf("Apellido %d", id),
	}

}

func checkError(e error) {

	defer func() {
		if e != nil {

			log.Printf("Write error: %v\n", e)

			log.Fatal("Entro en pánico. Error grave.")
		}
		log.Println("No errors found")
	}()

}

func panicRecover() {
	defer func() {
		if recovering := recover(); recovering != nil {
			println("Recuperamos la calma después de ", recovering)
		}
	}()
}
