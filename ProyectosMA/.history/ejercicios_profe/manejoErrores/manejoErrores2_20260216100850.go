package manejoErrores

import (
	"errors"
)

var erroNoEncontrado = errors.New("no encontrado")

const (
	pizza int = iota
	hamburgesa
)

var comida = map[int]string{
	1: "🍕",
	2: "🍔"}

func EjpMajenoErr() {

}
