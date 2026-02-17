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

	agenda = openAgenda()

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
	checkError(agenda.Close())

	checkError(err)

	for id, contact := range contacts {
		line := fmt.Sprintf("%d: %s %s\n", id, contact.name, contact.surname)
		n, err := agenda.WriteString(line)

		log.Printf("%d bytes written to %s file\n", n, agenda.Name())

		checkError(err)

	}
}

func createAgenda() *os.File {

	log.Println("Creating file...")

	path, error := os.Getwd()
	checkError(error)



	fichero, error := os.Create( fmt.Sprintf(("%s%s+agenda.txt",path,os.PathSeparator))

	func() {
		checkError(fichero.Close())
		log.Printf("File %s properly closed.\n", fichero.Name())
	}()

	log.Printf("File %s created.\n", fichero.Name())

	checkError(error)

	panicRecover()

	return fichero

}

func openAgenda() *os.File {
	agenda, err := os.OpenFile("agenda.txt", os.O_RDONLY, 0644)
	checkError(err)

	return agenda

}

func fillContactMap(contacts map[int]contactStruct) {
	log.Println("Filling contact map...")
	for i := 0; i < 10; i++ {
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
			panic("Entro en pánico. Error grave.")
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
