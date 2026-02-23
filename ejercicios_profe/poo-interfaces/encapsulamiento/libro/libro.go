package libro

import "fmt"

// Simulo una clase y sus atributos.
type Libro struct {
	titulo  string
	autor   string
	paginas int
}

/*
En cuanto a los cosntructores en Go, no hay un concepto de constctor tradicional como en otros lenguajes POO. En lugar de eso, se crea una función que devuelve una instancia de la estructura, inicializada con los valores necesarios.
*/

// Simulo un constructor lleno.
func NuevoLibro(titulo string, autor string, paginas int) *Libro {
	return &Libro{
		titulo:  titulo,
		autor:   autor,
		paginas: paginas,
	}
}

// Método set para establecer el título del libro.
func (l *Libro) SetTitulo(titulo string) {
	l.titulo = titulo
}

// Método get para obtener el título del libro.
func (l *Libro) GetTitulo() string {
	return l.titulo
}

// Método set para establecer el autor del libro.
func (l *Libro) SetAutor(autor string) {
	l.autor = autor
}

// Métodop get para obtener el autor del libro.
func (l *Libro) GetAutor() string {
	return l.autor
}

// Método set para establecer el número de páginas del libro.
func (l *Libro) SetPaginas(paginas int) {
	l.paginas = paginas
}

// Métodop get para obtener el número de páginas del libro.
func (l *Libro) GetPaginas() int {
	return l.paginas
}

// Método para imprimir la información del libro.
func (l *Libro) ImprimirInfo() {
	fmt.Printf("Titulo: %s\nAutor: %s\nN. Páginas: %d\n", l.titulo, l.autor, l.paginas)
}
