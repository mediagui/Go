package libro

import "fmt"

// Simulo una clase y sus atributos.
type Libro struct {
	Titulo  string
	Autor   string
	Paginas int
}

/*
En cuanto a los cosntructores en Go, no hay un concepto de constctor tradicional como en otros lenguajes POO. En lugar de eso, se crea una función que devuelve una instancia de la estructura, inicializada con los valores necesarios.
*/

// Simulo un constructor lleno.
func NuevoLibro(titulo string, autor string, paginas int) *Libro {
	return &Libro{
		Titulo:  titulo,
		Autor:   autor,
		Paginas: paginas,
	}
}

// Método para imprimir la información del libro.
func (l *Libro) ImprimirInfo() {
	fmt.Printf("Titulo: %s\nAutor: %s\nN. Páginas: %d\n", l.Titulo, l.Autor, l.Paginas)
}
