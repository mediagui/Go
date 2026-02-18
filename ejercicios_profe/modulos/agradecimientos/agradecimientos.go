package agradecimientos

import (
	"errors"
	"fmt"
	"math/rand"
)

// Función [Hola()], que recibe como parámetro un nombre y almacena en una variable un saludo a la persona específica.
func Hola(f_nombre string) (string, error) {
	/*
		[Sprint]: Se utiliza cuando se quiere guardar un texto en una variable.
		[Sprintf]: Se utiliza cuando se quiere guardar un texto formateado en una variable.
	*/

	// Compruebo que se esté recibiendo un nombre. En caso contrario, devuelvo un error.
	if f_nombre == "" {
		return "", errors.New("Nombre Vacío")
	}

	// En caso de que si que se haya recibido un nombre, lo devuelvo.
	mensaje := fmt.Sprintf(saludoRand(), f_nombre)

	// Devuelvo el saludo.
	return mensaje, nil
}

// Funcvión que recibe un slice de nombres y devuelve un mapa con los nombres y un posible error.
func HolaVarias(f_nombres []string) (map[string]string, error) {
	// Mapaa para asociar cada nombre con un mensaje.
	mensajes := make(map[string]string)

	// Itero sobre el los nombres, añadiendo cada saludo a la persona en el mapa.
	for _, nombre := range f_nombres {
		mensaje, err := Hola(nombre)
		if err != nil {
			return nil, err
		}
		mensajes[nombre] = mensaje
	}
	return mensajes, nil
}

// Función que devuelve uno de varios formatos de saludo. El mensaje jdevuelto se selecciona de forma aleatoria.
func saludoRand() string {
	// Declaro y defino un slice con "formatos" de mensajes.
	formatos := []string{
		"¡HOLA, %v! ¡Bienvenido/a!",
		"¡Qúe bueno verte, %v! (ironía)",
		"¡Saludos, %v! ¡Estarás encantado de conocerme!",
	}

	// Devuelvo un formato de mensaje seleccionado el índice de forma aleatoria.
	return formatos[rand.Intn(len(formatos))]
}
