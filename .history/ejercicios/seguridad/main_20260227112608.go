package main

// Las ctes en Go, por convención no se deben escribir en mayúsculas ya que
// las convertimos en públicas automaticamente
// para evitar el acceso desde fuera de la aplicación, habría que definirlas
// en un módulo llamado "internal"
const FILE_NAME = "auditoria.txt"

type Acceso struct {
	Id        uint
	Usuario   string
	Riesgo    int
	Categoria string
}

// Crea e Inicializa un array con 5 elementos de tipo Acceso
var eventos [5]Acceso = [5]Acceso{
	{Id: 1, Usuario: "admin", Riesgo: 1},
	{Id: 2, Usuario: "usuario1", Riesgo: 9},
	{Id: 3, Usuario: "usuario2", Riesgo: 2},
	{Id: 4, Usuario: "usuario3", Riesgo: 9},
	{Id: 5, Usuario: "usuario4", Riesgo: 4},
}

func main() {

}

func asignaCategorias(ev Acceso) Acceso {
	switch {
	case ev.Riesgo >= 1 && ev.Riesgo <= 3:
		ev.Categoria = "Seguro"
	case ev.Riesgo >= 4 && ev.Riesgo <= 6:
		ev.Categoria = "Usuario Avanzado"
	case ev.Riesgo >= 7 && ev.Riesgo <= 9:
		ev.Categoria = "Usuario Experto"
	default:
		ev.Categoria = "Usuario Desconocido"
	}
	return ev
}
