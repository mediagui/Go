package inout

import (
	"bufio"
	"fmt"
	"funciones/dto"
	"os"
	"strings"
)

const CAMPOS = "Nombre,Apellido1,Apellido2,Dni,Tfno"
var camposDisponibles []string = strings.Split(CAMPOS, ",")

var datoAPedir []string

func PedirInfo(user *dto.UsuarioStruct) {

	fmt.Printf("Campos disponibles: %v", camposDisponibles)
	fmt.Print("Introduzca los campos a pedir separados por comas: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		datoAPedir = []string{}
	} else {
		datoAPedir = strings.Split(input, ",")
	}
}
