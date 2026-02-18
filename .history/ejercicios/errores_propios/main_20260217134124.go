// Se pide desarrollar una aplicación que permita a los usuarios guardar y administrar contactos en un archivo.
// Hay que recordar manejar los posibles errores que puedan ocurrir.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

type contactStruct struct {
	id      int
	name    string
	surname string
}

func main() {

	log.Println("Starting...")

	idFlag := checkInputParameter()

	createAgenda()
	loadContents()
	showFileContents()

	if *idFlag != 0 {
		deleteRecordsWithId(*idFlag)
	}

	showFileContents()
	panicRecover()
	log.Println("End.")

	os.Exit(0)

}

func checkInputParameter() *int {

	idFlag := flag.Int("id", 0, "Contact ID to delete")
	flag.Parse()

	if *idFlag == 0 {
		log.Println("Id to delete no set. Loading data.")
	} else {
		log.Printf("Id to delete: %d", *idFlag)
	}
	return idFlag
}

func deleteRecordsWithId(i int) {

	file := openAgenda(true)

	fileSize := getFileSize(file)
	content := readFileContents(fileSize, file)
	file.Close()

	contentMap := make(map[int]contactStruct)

	// Parse content and populate contentMap
	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		var id int
		var contact contactStruct
		fmt.Sscanf(scanner.Text(), "%d: %s %s", &id, &contact.name, &contact.surname)
		if id != i {
			log.Printf("Adding contact: %v\n", contact)
			contentMap[id] = contact
		}
	}
	checkError(scanner.Err())

	file = openAgenda()
	defer file.Close()
	saveContactsToAgenda(contentMap)

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

func loadContents() {
	contacts := map[int]contactStruct{}
	fillContactMap(contacts)
	saveContactsToAgenda(contacts)
}

func showFileContents() {

	log.Println("Opening file for reading...")
	agendaContent, err := os.ReadFile("agenda.txt")
	checkError(err)

	log.Println("Showing file contents...")
	fmt.Printf("Contenido: \n%s\n", agendaContent)

}

func saveContactsToAgenda(contacts map[int]contactStruct) {
	var buffer bytes.Buffer
	for id, contact := range contacts {
		buffer.WriteString(fmt.Sprintf("%d: %s %s\n", id, contact.name, contact.surname))
	}

	file := openAgenda(false)
	_, err := file.WriteAt(buffer.Bytes(), 0)
	checkError(err)
}

func createAgenda() {

	log.Println("Creating file...")

	fichero, error := os.Create("agenda.txt")

	defer func() {
		checkError(fichero.Close())
		log.Printf("File %s properly closed.\n", fichero.Name())
	}()

	log.Printf("File %s created.\n", fichero.Name())

	checkError(error)

	panicRecover()

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
		id:      id,
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
