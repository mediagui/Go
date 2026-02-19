package main

import "libro"

func main() {
	l := libro.Libro{
		Nombre:  "El principito",
		Titulo:  "Un cuento",
		Paginas: 100,
	}

	l.PrintInfo()
}
