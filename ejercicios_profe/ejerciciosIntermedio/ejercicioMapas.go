package ejerciciosIntermedio

/*
###########
#Enunciado#
###########

Crear un mapa que tenga como claves los números del día de la semana y como valores el nombre del día de la semana correspondiente:

1 → Lunes
2 → Martes
3 → Miércoles
4 → Jueves
5 → Viernes
6 → Sábado
7 → Domingo


Luego, crea un buccle que pregunte al usuario si quiere modificar algún día de la semana y cuántos. Si es así, le preguntará el número del día y el nuevo nombre del día. Finalmente, imprimirá el diccionario modificado.
*/

import (
	"fmt"
)

var (
	cantModificar int
	clave         int
	valor         string
)

// Función que imprimirá los días de la semana.
func EjercicioMapas() {
	fmt.Println()
	fmt.Println(`
#################
#EJERCICIO MAPAS#
#################`)
	// Creo un diccionario con los días de la semana.
	semana := map[int]string{
		1: "Lunes",
		2: "Martes",
		3: "Miércoles",
		4: "Jueves",
		5: "Viernes",
		6: "Sábado",
		7: "Domingo",
	}
	// Imprimo el diccionario.
	fmt.Println(fondoBlanco, codigoAzul)
	// Utilizo un bucle for para imprimir el diccionario, para que se ordene de forma correcta.
	for i := 1; i <= len(semana); i++ {
		fmt.Println(i, "→", semana[i])
	}
	fmt.Println(reset)

	// Pido al usuario cuántos días desea modificar.
	fmt.Println("SISTEMA DE MODIFICACIÓN DE DÍAS DE LA SEMANA.")
	fmt.Println("Introduce cuántos días deseas modificar: ")
	fmt.Scan(&cantModificar)

	// Itero sobre la cantidad de días que el usuario desea modificar, preguntando la clave del dato que desea modificar y el nuevo valor.
	for i := 0; i < cantModificar; i++ {
		fmt.Print("Introduce la clave del dato que deseas modificar: ")
		fmt.Scan(&clave)
		fmt.Print("Introduce el nuevo valor: ")
		fmt.Scan(&valor)
		semana[clave] = valor
	}

	// Imprimo el diccionario modificado.
	fmt.Println(fondoNegro, codigoVerde)
	for i := 1; i <= len(semana); i++ {
		fmt.Println(i, "→", semana[i])
	}
	fmt.Println(reset)
}
