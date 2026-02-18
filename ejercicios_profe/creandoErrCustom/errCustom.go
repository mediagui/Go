package creandoerrcustom

import (
	"fmt"
)

// Función que devolverá la división de un número entre 0.
func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// return 0, errors.New("No se puede dividir entre 0, ssssstupido.")
		return 0, fmt.Errorf("No se puede imprimir, desde el paquete FMT, botarate")
	} else {
		return dividendo / divisor, nil
	}
}

func ErrCustom() {
	resultado, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("Resultado: ", resultado)
	}
}
