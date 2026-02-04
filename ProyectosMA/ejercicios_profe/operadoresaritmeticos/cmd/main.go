// Indico el paquete principal.
package main

// Importaciones.
import (
	"operadoresaritmeticos"
)

// Declaro la función principal [main], que actúa como punto de inicio de la ejecución del programa. No recibe parámetros de entrada ni retorna ningún valor
func main() {
	// Invoco a la función exportada del paquete [operaodres], para ejecutar la lógica de las funciones dentro de este.
	operadoresaritmeticos.MostrarOperaciones()
}
