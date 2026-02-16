package manejoErrores

import (
	"errors"
	"fmt"
	"strconv
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
	fmt.Print(buscar())
}

func buscar(clave string) (string, error){
	num, err := strConv.Atoi(clave)
	if err != nil{
		return "",fmt.Errorf("No se puede encontrar el elemento")
	}
}
