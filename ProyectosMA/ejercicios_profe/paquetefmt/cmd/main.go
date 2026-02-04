// Indico el paquete al que pertenece el archivo.
package main

// Importación de paquetes. Importo el [paquetefmt], para poder invocar a las funciones exportadas desde este punto.
import "paquetefmt"

// Defino la función principal. Esta actúa como punto de inicio de ejecución del programa. no recibe parámetros ni retorna valores.
func main() {
	// Llamo a la función  exportada [FuncionFmt()] para ejecutar la lógica de dicha función.
	paquetefmt.FuncionFmt()
}
