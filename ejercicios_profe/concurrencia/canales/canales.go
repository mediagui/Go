package concurrencia

import "fmt"

/*
#########
#CANALES#
#########

En go, los canales son una característica importante para la comunicación y sicronización entre [goroutines] dentro de un programa concurrente. Un canal es una estructura que permite enviar y recibir valores entre [goroutines], actuando como conducto a través del cual fluye la información.

#SINTAXIS#
canal := make(chan tipoDato)

Donde [tipoDato] especifica el tipo de datos que se enviarán a través del canal. Puede ser cualquier tipo de dato válido en Go como int, string, struct, etc
*/

func MostrarCanales() {

	// Creo un canal de tipo entero.
	ch := make(chan int)

	// Enviar valor a través del canal.
	ch <- 10

	// Recibir un valor del canal.
	valor := <-ch

	fmt.Println(valor)
}
