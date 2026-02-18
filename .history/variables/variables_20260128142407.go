package variables

import "fmt"

func GetUint8Value() uint8 {

	var value uint8 = 127

	return value
}

func TratoVariables() {

	var nombreUno string

	nombreUno = "Manolito1"

	fmt.Println(nombreUno)

	var nombreDos, apellido1 string
	nombreDos = "Manolito2"
	apellido1 = "Gafotas"

	fmt.Println(nombreDos, apellido1)

}
