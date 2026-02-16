package manejoErrores

import (
	"errors"
	"fmt"
	"strconv"
)

// Creo el error para un elemento no encontrado.
var errNoEncontrado = errors.New("no encontrado")

// Creo un mapa que contiene las comidas sobre las que voy a trabajar.
var comida = map[int]string{
	1: "34",
	2: "44",
}

// Función para comprobar si el elemento buscado se encuentra o no en el mapa.
func EjpManejoErr() {
	// Almaceno en una variable la búsqueda. Esta devuelve si se ha encontrado o no la comida y un posible error.
	encontrado, err := buscar("34")
	// Compruebo si se ha encontrado o no la comida.
	if errors.Is(err, errNoEncontrado) {
		fmt.Println("Se pudo manejar el error correctamente.")
		return
	}
	if err != nil {
		fmt.Println("Buscar()", err)
	}
	fmt.Println(encontrado)
}

// Función que realiza la búsqueda del código de un "emoji" de comida, utilizando una clave.
func buscar(f_clave string) (string, error) {
	// Convierto la clave recivida a número entero.
	num, err := strconv.Atoi(f_clave)
	if err != nil {
		// En caso de que falle la conversión, devuelvo el mensaje correspondiente.
		return "", fmt.Errorf("Strconv.Atoi(): %w", err)
	}

	// En caso de que encuentre el emoji correspondeinte, llamo a la función que busca la comida en el mapa.
	emoji, err := encontrarComida(num)
	if err != nil {
		return "", fmt.Errorf("encontrarComida(): %w", err)
	}

	// Devuevlo el "emoji" correspondiente y nil, para el error.
	return emoji, nil
}

// Función que busca el emoji de comida correspondiente a un identificador.
func encontrarComida(f_id int) (string, error) {
	// Compruebo si el identificador se enuentra en el mapa.
	valor, existe := comida[f_id]
	if !existe {
		// Si no existe, devuellvo un error personalizado.
		return "", errNoEncontrado
	}

	// En caso de que si que exista, devuelvo el "emoji" y nil como error.
	return valor, nil

}
