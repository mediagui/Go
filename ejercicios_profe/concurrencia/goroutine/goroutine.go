package concurrencia

/*
##############
#CONCURRENCIA#
##############

La concurrencia es importante puesto que nos ayuda a aprovechar de forma efectiva los recursos de hardware y mejorar el rendimiento de nuestras aplicacines.

Los canales son herramientas para la comunicación y sincronización entre las [goroutines]. Nos permiten enviar y recibir datos de manera segura y coordinada, evitando condiciones de carrera y garantizando una ejecucuión coherente de nuestras rutinas (goroutines)

La concurrencia es una característica del lenguaje que permite la ejecución de múltiples tareas (gorutines, tareas similares a hilos). So independientes entre sí y se ejecutan de manera concurrente, esto quiere decir que se pueden ejecutar de forma simultánea o en un orden no determinista.
*/

import (
	"fmt"
	"net/http"
	"time"
)

func MostrarGoroutine() {
	inicio := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}

	// Bucle para verificar el estado de cada api.
	for _, api := range apis {
		// Añado una [goroutine], para ello es necesaria la palabra [go]
		go comprobarAPI(api)
	}

	// Hago que el programa espere x segundos antes de finalizar.
	time.Sleep(2 * time.Second)

	tiempoFinal := time.Since(inicio)
	fmt.Printf("¡LISTO! Ha tardado %v segundos\n", tiempoFinal.Seconds())
}

func comprobarAPI(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("Error: ¡%s está caído!\n", api)
		return
	}

	fmt.Printf("EXITO ¡%s está en funcionamiento!\n", api)
}

/*

Sin concurrencia.

1. Se importan los paquetes necesarios:

	"fmt" para imprimir mensajes en la consola, "net/http" para realizar solicitudes HTTP y "time" para medir el tiempo de ejecución.

2. En la función [main()]:

	se registra el tiempo de inicio de la ejecución

3. Se define el slice que contiene las URLs a las APIs a verifficar.

4. A continuación se inicia un bucle for para recorrer cada una de las URLs en apis

5. En cada iteración del bucle, se llama al afunción que comprueba cada API utilizando una solicitud HTTP GET a la URL espcificada por api, utilizando http.Get(api). La respuesta obtenida y un posible error se asignan a las variables [_] y [err] respectivamente.

Si se produce un error durante la solicitud, se imprime un mensaje de error indicando que la API está caída.

En caso contrario, se imprime un mensaje de éxito indicando que la API está en funcionamiento.

Una vez verificadas todas las APIs, se calcula el tiempo transcurrido.

Finalmente se imprime un mensaje indicando el tiempo total de ejecución en segundos.
*/
