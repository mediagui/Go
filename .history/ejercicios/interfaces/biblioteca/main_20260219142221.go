package main

import "biblioteca/libro"

func main() {
	libro := libro.Libro{
		Nombre:  "El principito",
		Titulo:  "Un cuento",
		Paginas: 100,
	}

	libro.PrintInfo()

	libro2 := NewLibro("Cien años de soledad", "Realismo mágico", 400)
	libro2.PrintInfo()

}
