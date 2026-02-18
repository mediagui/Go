package inout

import (
	"bufio"
	"fmt"
	"funciones/dto"
	"os"
	"strings"
)

// Campos disponibles para pedir
const CAMPOS = "Nombre,Apellido1,Apellido2,Dni,Tfno"

var camposDisponibles []string = strings.Split(CAMPOS, ",")

var datoAPedir []string

func PedirInfo(user *dto.UsuarioStruct) {

	// fmt.Printf("Campos disponibles: %v", camposDisponibles)
	fmt.Println("Introduzca la información pedida. ")

	reader := bufio.NewReader(os.Stdin)

	//Pedimos todos los datos definidos en la constante CAMPOS a través del slice
	for _, v := range camposDisponibles {

		fmt.Printf("%s: ", v)              //Pedimos el valor
		input, _ := reader.ReadString('\n') //Leemos el dato
		input = strings.TrimSpace(input)    //Devuelve la cadena eliminando los espacios al ppio y al final de cada elemento.

		dto.UsuarioStruct[v] = input //Almacenamos en el struct

	}

	// func PintarInfo( user dto.dto.UsuarioStruct){

	// 	fmt.Printf("Campos disponibles: %v", camposDisponibles)
	// 	fmt.Print("Introduzca los campos a pedir separados por comas: ")

	// 	input = strings.TrimSpace(input) //Devuelve la cadena eliminando los espacios al ppio y al final de cada elemento.
	// 	if input == "" {
	// 		//Si pulsamos intro, pedimos todos los campos
	// 		datoAPedir = camposDisponibles
	// 	} else {
	// 		datoAPedir = strings.Split(input, ",")

	// }
}
