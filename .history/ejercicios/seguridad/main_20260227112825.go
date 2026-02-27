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

func inicializaArray(ev []Acceso) {

	for i, _ := range ev {
		ev[i].Id = uint(i + 1)
		ev[i].Usuario = "usuario" + string(rune(i+1+'0'))
		ev[i].Riesgo = 0
		ev[i].Categoria = "Sin Categorizar"
	}
}

func asignaCategorias(ev Acceso) Acceso {
	switch {
	case ev.Riesgo >= 1 && ev.Riesgo <= 3:
		ev.Categoria = "Seguro"
	case ev.Riesgo >= 4 && ev.Riesgo <= 7:
		ev.Categoria = "Advertencia"
	case ev.Riesgo >= 8 && ev.Riesgo <= 10:
		ev.Categoria = "ATAQUE CRÍTICO"
	default:
		ev.Categoria = "Nivel Desconocido"
	}
	return ev
}
