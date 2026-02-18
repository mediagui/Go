// Se pide desarrollar una aplicación que permita a los usuarios guardar y administrar contactos en un archivo.
// Hay que recordar manejar los posibles errores que puedan ocurrir.
package main

import (
	"bufio"
	"bytes"
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

	agenda := loadContents()
	showFileContents(agenda)

	deleteRecordsWithId(3)

	showFileContents(agenda)

	panicRecover()

	log.Println("End.")

	os.Exit(0)

}

func deleteRecordsWithId(i int) {
	file := openAgenda(false)
	defer file.Close()

	fileSize := getFileSize(file)
	content := readFileContents(fileSize, file)

	contentMap := make(map[int]contactStruct)

	// Parse content and populate contentMap
	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		var id int
		var contact contactStruct
		fmt.Sscanf(scanner.Text(), "%d: %s %s", &id, &contact.name, &contact.surname)
		if id != i {
			contentMap[id] = contact
		}
	}
	checkError(scanner.Err())

	saveContactsToAgenda(contentMap, file)

}

func readFileContents(fileSize int64, fichero *os.File) []byte {
	var fileContents []byte = make([]byte, fileSize)
	bytesRead, err := fichero.Read(fileContents)
	log.Printf("Bytes read %d\n", bytesRead)
	checkError(err)

	return fileContents
}

func getFileSize(fichero *os.File) int64 {
	fileStats, err := fichero.Stat()
	checkError(err)
	return fileStats.Size()
}

func loadContents() *os.File {
	contacts := map[int]contactStruct{}
	fillContactMap(contacts)

	agenda := createAgenda()
	saveContactsToAgenda(contacts, agenda)
	return agenda
}

func showFileContents(agenda *os.File) {
	log.Println("Opening file for reading...")
	agenda = openAgenda()

	log.Println("Showing file contents...")
	showContentOf(agenda)
}

func showContentOf(agenda *os.File) {

	agendaContent, err := os.ReadFile(agenda.Name())
	checkError(err)
	fmt.Printf("Contenido: \n%s\n", agendaContent)

}

func saveContactsToAgenda(contacts map[int]contactStruct, agenda *os.File) {

	agenda, err := os.OpenFile(agenda.Name(), os.O_APPEND|os.O_WRONLY, 0644)
	checkError(err)

	defer agenda.Close()

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

	defer func() {
		checkError(fichero.Close())
		log.Printf("File %s properly closed.\n", fichero.Name())
	}()

	log.Printf("File %s created.\n", fichero.Name())

	checkError(error)

	panicRecover()

	return fichero

}

// Simulamos parametros opcionales
func openAgenda(asReadOnly ...bool) *os.File {
	var perm int

	_readOnly := len(asReadOnly) > 0 && asReadOnly[0]

	if _readOnly {
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

	if e != nil {

		log.Printf("Write error: %v\n", e)

		log.Fatal("Entro en pánico. Error grave.")
	}
	log.Println("No errors found")

}

func panicRecover() {
	defer func() {
		if recovering := recover(); recovering != nil {
			println("Recuperamos la calma después de ", recovering)
		}
	}()
}
