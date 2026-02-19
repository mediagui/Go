package libro

type Libro struct {
	Nombre  string
	Titulo  string
	Paginas int
}

func (l Libro) NewLibro(s string, param2 string, i int) any {
	panic("unimplemented")
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
