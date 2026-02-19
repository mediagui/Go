package libro

type Libro struct {
	Nombre  string
	Titulo  string
	Paginas int
}

func (l Libro) PrintInfo() {
	println(l.Nombre, l.Titulo, l.Paginas)
}

// Constructor para Libro que devuelva una referencia al objeto creado
func NewLibro(nombre, titulo string, paginas int) *Libro {
	return &Libro{
		Nombre:  nombre,
		Titulo:  titulo,
		Paginas: paginas,
	}
}
