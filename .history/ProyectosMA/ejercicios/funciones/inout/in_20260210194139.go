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

type tipoMetodo string

const (
	SET tipoMetodo = "Set"
	GET tipoMetodo = "Get"
)

var camposDisponibles []string = strings.Split(CAMPOS, ",")

var datoAPedir []string

// Método para guardar el valor del campo pedido
// campo: nombre del campo
// valor: valor del campo
// user: puntero al dto que almacenará los valores
func guardaCampoEn(campo string, valor string, userDto *dto.UsuarioStruct) {

	// Obtenemos el puntero (usando reflect)
	userDtoValue := reflect.ValueOf(userDto)

	// Construimos el setter (e.g., "SetNombre" a partir de "Nombre")
	nombreMetodo, metodo := metodoAInvocar(campo, userDtoValue, SET)

	// Guardamos el valor en el dto
	invocaElMetodoEnDto(metodo, valor, nombreMetodo)

}

// Almacena el valor introducido en el dto utilizando el método generado
func invocaElMetodoEnDto(metodo reflect.Value, v string, nombreMetodo string) {

	// Si el método existe...
	if metodo.IsValid() {
		log.Printf("Método/Func [%s] encontrado.\n", nombreMetodo)
		// ... lo invocamos y guardamos el valor o devolvemos el valor almacenado
		if v != "" {
			// Set method call
			metodo.Call([]reflect.Value{reflect.ValueOf(v)})
		} else {
			// Get method call
			resultado := metodo.Call([]reflect.Value{})
			if len(resultado) > 0 {
				log.Printf("Valor obtenido de [%s]: %v\n", nombreMetodo, resultado[0].Interface())
			}
		}

		log.Printf("Valor [%s] almacenado en [%s].\n", v, nombreMetodo)

	} else {
		fmt.Printf("Método/Func [%s] inválido\n", nombreMetodo)
	}
}

func metodoAInvocar(k string, userDtoValue reflect.Value, tipo tipoMetodo) (string, reflect.Value) {

	log.Println("Generando el método a invocar...")

	//Aquí no es de mucha utilidad, pero muestra como usar tipos de datos personalizados
	methodName := string(tipo) + k

	// Obtenemos el método
	method := userDtoValue.MethodByName(methodName)

	log.Printf("Método a invocar: [%s].", methodName)

	return methodName, method
}

// Muestra el menú principal para elegir entre introducir o mostrar datos
func MuestraMenuPrincipal() {
	fmt.Println("1. Introduzca datos\n2. Seleccione campos a mostrar\n\tElija: ")
}

// Muestra el menú para introducir los campos que se desean mostrar
func MuestraMenuVisualizacionDatos() {
	fmt.Println("Los posibles campos a mostrar son [Nombre, Apellido1, Tfno y Dni]\n\tEscriba el nombre de los campos separados por coma: ")
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
	fmt.Scanln(&campos)
	sliceCampos := strings.Split(campos, ",")

	// Eliminar los espacios en blanco de cada uno de los elementos del sliceCampos
	for i := range sliceCampos {
		sliceCampos[i] = strings.TrimSpace(sliceCampos[i])
	}

	return sliceCampos

}

func PedirCamposAGuardar(user *dto.UsuarioStruct) {

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

func GeValuesFromDto(fields []string, userDto *dto.UsuarioStruct) {

	//var dto dto.UsuarioStruct

	// Obtenemos el puntero (usando reflect)
	userDtoValue := reflect.ValueOf(userDto)

	for i := 0; i < len(fields); i++ {

		nombreMetodo, metodo := metodoAInvocar(fields[i], userDtoValue, GET)

		invocaElMetodoEnDto(metodo, "", nombreMetodo)
	}
	// Construimos el setter (e.g., "SetNombre" a partir de "Nombre")

}
