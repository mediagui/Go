package libro

type Libro struct {
	Nombre  string
	Titulo  string
	Paginas int
}

func (l Libro) PrintInfo() {
	println(l.Nombre, l.Titulo, l.Paginas)
}
