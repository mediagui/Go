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

// Muestra el menú principal para elegir entre introducir o mostrar datos
func MuestraMenuPrincipal() {
	fmt.Println("1. Introduzca datos\n2. Seleccione campos a mostrar\n\tElija: ")
}

// Muestra el menú para introducir los campos que se desean mostrar
func MuestraMenuVisualizacionDatos() {
	fmt.Println("Los posibles campos a mostrar son: \n Nombre, Apellidos, Tfno y Dni \n Escriba los campos que quiere mostrar: ")
}

// Lee en consola la opción seleccionada y la devuelve
func GetOpcionSeleccionada() int {
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}

// Lee por consola los campos a mostrar y devuelve un slice con el nombre de los mismos.
func GetCamposAMostrar() []string {

	var campos string
	fmt.Sscanln(campos)
	sliceCampos := strings.Split(strings.TrimSpace(campos), ",")

	// Eliminar los espacios en blanco de cada uno de los elementos del sliceCampos
	for i := range sliceCampos {
		sliceCampos[i] = strings.TrimSpace(sliceCampos[i])
	}

	return sliceCampos

}
