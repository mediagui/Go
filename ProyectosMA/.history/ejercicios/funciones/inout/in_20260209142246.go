package inout

import (
	"bufio"
	"fmt"
	"funciones/dto"
	"log"
	"os"
	"reflect"
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
	for _, campoAGuardar := range camposDisponibles {

		log.Printf("Pidiendo campo %s.\n", campoAGuardar)

		fmt.Printf("%s: ", campoAGuardar)                //Pedimos el valor
		valorAGuardar, _ := reader.ReadString('\n')      //Leemos el dato
		valorAGuardar = strings.TrimSpace(valorAGuardar) //Devuelve la cadena eliminando los espacios al ppio y al final de cada elemento.

		guardaCampoEn(campoAGuardar, valorAGuardar, user)

	}
}

// Método para guardar el valor del campo pedido
// campo: nombre del campo
// valor: valor del campo
// user: puntero al dto que almacenará los valores
func guardaCampoEn(campo string, valor string, user *dto.UsuarioStruct) {

	// Obtenemos el puntero (usando reflect)
	userValue := reflect.ValueOf(user)

	// Construimos el setter (e.g., "SetNombre" a partir de "Nombre")
	nombreMetodo, metodo := metodoAInvocar(campo, userValue)

	// Guardamos el valor en el dto
	guardaElValorEnDto(metodo, valor, nombreMetodo)

}

func guardaElValorEnDto(metodo reflect.Value, v string, nombreMetodo string) {

	// Si el método existe...
	if metodo.IsValid() {
		log.Printf("Método/Func [%s] encontrado.\n", nombreMetodo)
		// ... lo invocamos y guardamos el valor
		metodo.Call([]reflect.Value{reflect.ValueOf(v)})
		log.Printf("Valor [%s] almacenado en [%s].\n", v, nombreMetodo)

	} else {
		fmt.Printf("Método/Func [%s] inválido\n", nombreMetodo)
	}
}

func metodoAInvocar(k string, userValue reflect.Value) (string, reflect.Value) {

	log.Println("Generando el método a invocar...")

	methodName := "Set" + k

	// Obtenemos el método
	method := userValue.MethodByName(methodName)

	log.Printf("Método a invocar: [%s].", methodName)

	return methodName, method
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
