package main

import (
	"fmt"

	"github.com/nox0V0saw/agradecimientos"
)

func main() {
	mensaje, err := agradecimientos.Hola("Víctor")

	if err != nil {
		fmt.Println("A ocurrido un error:", err)
		return
	}

	fmt.Println(mensaje)
}
