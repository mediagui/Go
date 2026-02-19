# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalación

Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/nox0V0saw/agradecimientos
```

## Uso

Ejemplo de cómo utilizar el paquete en nuestro código:

```go
package main

import (
  "fmt"
  "github.com/nox0V0saw/agradecimientos"
)

func main(){
  mensaje, err := agradecimientos.Hola("Víctor")

  if err != nil {
    fmt.Println("A ocurrido un error:",err)
    return
  }

  fmt.Println(mensaje)
}
```
Este ejemplo importa el paquete [github.com/nox0V0saw/agradecimientos] y llama a la función [Hola()] para obtener un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.