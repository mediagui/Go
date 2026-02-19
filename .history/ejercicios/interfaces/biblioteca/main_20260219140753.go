package main

import "fmt"

type Libro struct {
	Nombre  string
	Titulo  string
	Paginas int
}

func (l Libro) PrintInfo() {
	fmt.Printf("%s - %s (%d páginas)\n", l.Nombre, l.Titulo, l.Paginas)
}

func main() {
	libro := Libro{
		Nombre:  "El principito",
		Titulo:  "Un cuento",
		Paginas: 100,
	}

	libro.PrintInfo()
}
