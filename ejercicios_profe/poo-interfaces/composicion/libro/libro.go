package libro

import (
	"fmt"
)

// Simulo una clase y sus atributos.
type Libro struct {
	titulo  string
	autor   string
	paginas int
}

/*
La composición es un mecanismo por el cual un tipo contiene a otro dentro de sí, reutilizando su comportamiento.
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

// Creo un nuevo tipo de dato, que heredará del libro original.
type LibroTexto struct {
	Libro
	editorial string
	nivel     string
}

// Constructor para los libros de texto.
func NuevoLibroTexto(titulo, autor string, paginas int, editorial, nivel string) *LibroTexto {
	return &LibroTexto{
		Libro:     Libro{titulo, autor, paginas},
		editorial: editorial,
		nivel:     nivel,
	}
}

// Método para imprimir la información del libro.
func (l *LibroTexto) ImprimirInfo() {
	fmt.Printf("Titulo: %s\nAutor: %s\nN. Páginas: %d\nEditorial: %s\nNivel: %s\n",
		l.titulo, l.autor, l.paginas, l.editorial, l.nivel)
}
