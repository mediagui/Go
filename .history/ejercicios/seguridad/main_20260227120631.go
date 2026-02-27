package main

import "fmt"

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
const NIVEL_DESCONOCIDO TipoEvento = "Nivel Desconocido"

type Acceso struct {
	Id        uint
	Usuario   string
	Riesgo    int
	Categoria TipoEvento
}

// Crea e Inicializa un array con 5 elementos de tipo Acceso
var eventos [5]Acceso

func main() {
	eventos = inicializaArray(eventos)
	eventos = asignaCategorias(eventos)

	alertaRiesgoCritico(eventos)

	fmt.Printf("Media de riesgo: %v\n", calculaMediaDeRiesgo(eventos))

	pintaArray()

}

// Inicializamos el array
// Podría usar punteros
func inicializaArray(ev [5]Acceso) [5]Acceso {

	for i := 0; i < 5; i++ {
		ev[i].Id = uint(i + 1)
		ev[i].Usuario = "usuario" + string(rune(i+1+'0'))
		ev[i].Riesgo = (i + 1) * 2
		ev[i].Categoria = "Sin Categorizar"
	}
	return ev
}

// Asigna las categorías definicas
func asignaCategorias(ev [5]Acceso) [5]Acceso {

	for e := range ev {

		ev[e].Categoria = obtieneCategoriaRiesgo(ev[e])
	}
	return ev
}

func obtieneCategoriaRiesgo(ev Acceso) TipoEvento {
	var nivelEvento TipoEvento

	switch {
	case ev.Riesgo >= 1 && ev.Riesgo <= 3:
		nivelEvento = SEGURO
	case ev.Riesgo >= 4 && ev.Riesgo <= 7:
		nivelEvento = ADVERTENCIA
	case ev.Riesgo >= 8 && ev.Riesgo <= 10:
		nivelEvento = ATAQUE_CRITICO
	default:
		nivelEvento = NIVEL_DESCONOCIDO
	}
	return nivelEvento
}

func calculaMediaDeRiesgo(ev [5]Acceso) TipoEvento {
	var totalRiesgo int
	for i := range 5 {
		totalRiesgo += ev[i].Riesgo
	}
	media := totalRiesgo / 5
	return obtieneCategoriaRiesgo(Acceso{Riesgo: media})

}

func alertaRiesgoCritico(ev [5]Acceso) {
	for i := range 5 {
		if ev[i].Categoria == ATAQUE_CRITICO {
			escribeEnConsola("ALERTA", func(ev [5]Acceso) TipoEvento {
				return ev[i].Categoria
			})
		}
	}
}

func escribeEnConsola(descripcion string, f func([5]Acceso) TipoEvento) {
	resultado := f(eventos)
	fmt.Printf("%s: %v\n", descripcion, resultado)
}

// Pinta en consola el contenido del array
func pintaArray() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Evento %d: %v\n", i+1, eventos[i])
	}
}
