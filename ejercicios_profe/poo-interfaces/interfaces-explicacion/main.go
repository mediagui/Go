package main

// Importaciones.
import (
	"fmt"
	// Importo la librería para crear los objetos.
	"interfaces-animal/animal"
)

// Función principal.
func main() {
	// Creo los objetos de los animales correspondientes.
	miPerro := animal.Perro{Nombre: "Perla"}
	miGato := animal.Gato{Nombre: "Vinagrito"}

	// Gracias a la interface creada, podemos pasar ambos tipos a la misma función, que se encargará de realizar el sonido correspondiente.
	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	fmt.Println(`
	Con lista de animales
	`)
	// Creo una lista con distintos animales.
	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Luna"},
		&animal.Perro{Nombre: "Menox"},
		&animal.Gato{Nombre: "Figaro"},
	}

	// Itero sobre la lista de animales y llamo la función  [Sonido()] de cada uno.
	for _, animal := range animales {
		animal.Sonido()
	}
}
