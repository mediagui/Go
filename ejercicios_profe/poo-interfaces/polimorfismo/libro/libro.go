package libro

import (
	"fmt"
)

/*
##############
#POLIMORFISMO#
##############

El polimorfismo es la capacidad de los objetos de diferentes clases para responder al mismo mensaje o método. En Go, se puede implementar el polimorfismo a través de interfaces.

Una interfaz es un conjunto de métodos que una estructura debe implementar para satisfacer la interfaz. Esto permite que una estructura implemente múltiples interfaces y, por lo tanto, tenga diferentes comportamientos en diferentes contextos.

En Go, el polimorfismo se puede lograr mediante interfaces. Podemos definir una interfaz común para que tanto la estructura Libro como la estructura [LibroTexto] implemente. Luego, podríamos crear una función para que acepte un parámetro de esta interfaz y pueda manejar ambos tipos de estructuras.

¿Por qué es útil?

El polimorfismo permite:

	1. Escribir código más flexible.
		Una función puede trabajar con distintos tipos sin saber cuáles son exactamente.

	2. Reducir duplicaciones.
		No necesitamos escribir una función distinta para cada tipo.

	3. Desacoplar el código.
		Tu función depende de un comportamiento, no de un tipo concreto.

	4. Cambiar implementaciones sin romper nada.
		Podemos sustituir un tipo por otro mientras cumpla la misma interfaz.
*/

// Simulo una clase y sus atributos.
type Libro struct {
	titulo  string
	autor   string
	paginas int
}

/*
La composición es un mecanismo por el cual un tipo contiene a otro dentro de sí, reutilizando su comportamiento.
*/

type Imprimible interface {
	ImprimirInfo()
}

// Para manejar la interface, podemos crear una función que ejecute la acción.
func Imprimir(I Imprimible) {
	I.ImprimirInfo()
}

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
