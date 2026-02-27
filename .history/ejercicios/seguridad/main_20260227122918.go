package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Las ctes en Go, por convención no se deben escribir en mayúsculas ya que
// las convertimos en públicas automaticamente
// para evitar el acceso desde fuera de la aplicación, habría que definirlas
// en un módulo llamado "internal"
const NOMBRE_ARCHIVO = "auditoria.txt"

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
	Hora      time.Time
}

// Crea e Inicializa un array con 5 elementos de tipo Acceso
var eventos [5]Acceso

func main() {

	wg := sync.WaitGroup{}

	eventos = inicializaArray(eventos)
	eventos = asignaCategorias(eventos)

	alertaRiesgoCritico(eventos)

	escribeEnConsola("Media de riesgo", func(ev [5]Acceso) TipoEvento {
		return calculaMediaDeRiesgo(ev)
	})

	guardaEnArchivo(eventos)

}

// Inicializamos el array
// Podría usar punteros
func inicializaArray(ev [5]Acceso) [5]Acceso {

	for i := 0; i < 5; i++ {
		ev[i].Id = uint(i + 1)
		ev[i].Usuario = "usuario" + string(rune(i+1+'0'))
		ev[i].Riesgo = (i + 1) * 2
		ev[i].Categoria = "Sin Categorizar"
		ev[i].Hora = time.Now()
	}

	log.Println("Array inicializado")

	return ev
}

// Asigna las categorías definicas
func asignaCategorias(ev [5]Acceso) [5]Acceso {

	for e := range ev {

		ev[e].Categoria = obtieneCategoriaRiesgo(ev[e])
	}

	log.Println("Categorías asignadas")

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

	log.Println("Categoría asignada:", nivelEvento)

	return nivelEvento
}

func calculaMediaDeRiesgo(ev [5]Acceso) TipoEvento {

	log.Println("Calculando media de riesgo")

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

func guardaEnArchivo(ev [5]Acceso) {

	f, err := os.OpenFile(NOMBRE_ARCHIVO, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Error al crear o acceder al fichero", NOMBRE_ARCHIVO)
		return
	}
	defer f.Close()

	for i := range 5 {

		registro := fmt.Sprintf(
			"Hora %v Evento %d: %v Riesgo %d\n",
			ev[i].Hora.Format("15:04:05"),
			i+1,
			ev[i].Categoria,
			ev[i].Riesgo,
		)

		f.WriteString(registro)
		log.Printf("Guardando registro %s", registro)
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
