package concurrencia

import (
	"fmt"
	"net/http"
	"time"
)

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

	inicio := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}

	// Canal para almacenar cada respuesta.
	ch := make(chan string)
	// Bucle para verificar el estado de cada api.
	for _, api := range apis {
		// Añado una [goroutine], para ello es necesaria la palabra [go]
		go comprobarAPI(api, ch)
		// fmt.Printf(<-ch)
	}

	// Imprimo el contenido del canal
	// fmt.Printf(<-ch)
	// fmt.Printf(<-ch)
	// fmt.Printf(<-ch)
	// fmt.Printf(<-ch)
	// fmt.Printf(<-ch)
	// fmt.Printf(<-ch)

	fmt.Println(`
	Con bucle for
	`)

	// Utilizo un bucle for para mostrar todas las apis.
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	// Hago que el programa espere x segundos antes de finalizar.
	time.Sleep(2 * time.Second)

	tiempoFinal := time.Since(inicio)
	fmt.Printf("¡LISTO! Ha tardado %v segundos\n", tiempoFinal.Seconds())
}

func comprobarAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Error: ¡%s está caído!\n", api)
		return
	}

	ch <- fmt.Sprintf("EXITO ¡%s está en funcionamiento!\n", api)
}
