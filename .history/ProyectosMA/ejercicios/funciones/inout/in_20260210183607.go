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
