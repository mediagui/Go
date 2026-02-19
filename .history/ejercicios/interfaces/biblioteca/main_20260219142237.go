package main

import "biblioteca/libro"

func main() {
	libro1 := libro.Libro{
		Nombre:  "El principito",
		Titulo:  "Un cuento",
		Paginas: 100,
	}

	libro1.PrintInfo()

	libro2 := libro.NewLibro("Cien años de soledad", "Realismo mágico", 400)
	libro2.PrintInfo()

}
