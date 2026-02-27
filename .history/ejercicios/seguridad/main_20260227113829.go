package main

// Las ctes en Go, por convención no se deben escribir en mayúsculas ya que
// las convertimos en públicas automaticamente
// para evitar el acceso desde fuera de la aplicación, habría que definirlas
// en un módulo llamado "internal"
const FILE_NAME = "auditoria.txt"

// Definición de tipo "custom"
type TipoEvento string

const SEGURO TipoEvento = "Seguro"
const ADVERTENCIA TipoEvento = "Advertencia"
const ATAQUE_CRITICO TipoEvento = "ATAQUE CRÍTICO"

type Acceso struct {
	Id        uint
	Usuario   string
	Riesgo    int
	Categoria TipoEvento
}

// Crea e Inicializa un array con 5 elementos de tipo Acceso
var eventos [5]Acceso = make([5]Acceso, 5)

func main() {
	inicializaArray(eventos[:])
}

// Inicializamos el array
// Podría usar punteros
func inicializaArray(ev []Acceso) []Acceso {

	for i := range ev {
		ev[i].Id = uint(i + 1)
		ev[i].Usuario = "usuario" + string(rune(i+1+'0'))
		ev[i].Riesgo = i * 2
		ev[i].Categoria = "Sin Categorizar"
	}
	return ev
}

// Asigna las categorías definicas
func asignaCategorias(ev []Acceso) []Acceso {

	for e := range ev {
		switch {
		case ev[e].Riesgo >= 1 && ev[e].Riesgo <= 3:
			ev[e].Categoria = SEGURO
		case ev[e].Riesgo >= 4 && ev[e].Riesgo <= 7:
			ev[e].Categoria = ADVERTENCIA
		case ev[e].Riesgo >= 8 && ev[e].Riesgo <= 10:
			ev[e].Categoria = ATAQUE_CRITICO
		default:
			ev[e].Categoria = "Nivel Desconocido"
		}
	}
	return ev
}
