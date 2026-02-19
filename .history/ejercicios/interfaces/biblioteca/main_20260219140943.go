package main

import "libro"

func main() {
	libro := libro.Libro{
		Nombre:  "El principito",
		Titulo:  "Un cuento",
		Paginas: 100,
	}

	libro.PrintInfo()
}
