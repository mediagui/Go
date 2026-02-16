package manejoErrores2

import (
	"errors"
	"fmt"
	"strconv"
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
	fmt.Print(encontrado)
}

func buscar(clave string) (string, error) {
	num, err := strconv.Atoi(clave)
	if err != nil {
		return "", fmt.Errorf("No se puede encontrar el elemento")
	}
	emoji, err := encontrarComida(tipoComida(num))
	if err != nil {
		return "", errors.New("No se encuentra el elemento")
	}
	return emoji, nil
}

func encontrarComida(id tipoComida) (string, error) {

	valor, existe := comida[id]
	if !existe {
		return "", erroNoEncontrado
	}
	return valor, nil
}
