package main

import (
	// concurrencia "concurrencia/goroutine"
	canales "concurrencia/canales"
)

// Importanciones

func main() {
	// concurrencia.MostrarGoroutine()
	canales.MostrarCanales()
}

/*
############
#EN RESUMEN#
############

Las goroutines son funciones independientes que se ejecutan de forma concurrente y ligera dentro de un programa Go. Hemos visto cómo las goroutines son útiles para realizar múltiples tareas simultáneamente.

Los canales, por otro lado, son herramientas para la comunicación y sincornización entre goroutines.  Hemos visto cómo enviar y recibir datos de manera segura y coordinada utilizando canales.
*/
