package main

// Importaciones.
import (
	// Importo la librería para crear los objetos.
	"biblioteca/libro"
	"fmt"
)

// Función principal.
func main() {
	miLibro := libro.NuevoLibro("El Señor de las Moscas", "William Golding", 288)

	// Creo un nuevo libro para ejemplificar el concepto de encapsulamiento.
	miLibroDos := libro.NuevoLibro("El Diario de Ana Frank", "Ana Frank", 1945)

	miLibro.ImprimirInfo()
	miLibroDos.ImprimirInfo()

	fmt.Println(`##################
	#Con Polimorfismo#
	##################`)
	libro.Imprimir(miLibro)
	libro.Imprimir(miLibroDos)
}
