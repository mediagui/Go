module ProyectosGo

// Aquí llegamos escribiendo en la consola [go mod init nombrePaquete], en este caso, [go mod init ProyectosGo]

// Para importar un paquete, utilizo el comando [go get nombrepaquete]. En este caso voy a importar el paquete [quote], dentro del paquete [rsc.io], que me permite, entre otras cosas, imprimir el mítico mensaje de "hola mundo". Como vemos, también se ha creado un archivo [go.sum]

go 1.25.6

// Como vemos, una vez ejecutado el "manejador de módulos, se nos han escrito una serie de dependencias. Ahora, vamos a importar el paquete en nuestro archivo principal.
require (
	golang.org/x/text v0.33.0 // indirect
	rsc.io/quote v1.5.2
	rsc.io/sampler v1.99.99 // indirect
)

// Paquetes para los colores de Github
require (
	github.com/fatih/color v1.18.0
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.40.0 // indirect
)
