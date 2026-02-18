package manejoErrores

import (
	"errors"
	"fmt"
)

var erroNoEncontrado = errors.New("no encontrado")

type tipoComida int

const (
	pizza tipoComida = iota
	hamburgesa
)

var comida = map[tipoComida]string{
	1: "🍕",
	2: "🍔"}

func EjpMajenoErr() {

	encontrado, err := buscar("34")

	if errors.Is(err, erroNoEncontrado) {
		fmt.Println("No se pudo manejar el error correctamente")
	}
	if err != nil {
		fmt.Println("Buscar()", err)
	}

}
