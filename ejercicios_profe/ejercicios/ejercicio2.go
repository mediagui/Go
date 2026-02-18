// Indco el paquete al que pertenece el archivo.
package ejercicios

// Importación.
import "fmt"

// Decalro y defino la constante [IVA]
const IVA = 0.21

/*
Función [ImprimirFctura()]. Esta imprimirá los datos de la factura, según las variables y constantes indicadas.
*/
func ImprimirFctura() {
	// Variables.
	p1Nombre, p1Precio := "Monitor 24\"", 149.99
	p2Nombre, p2Precio := "Teclado Mecánico", 45.50
	p3Nombre, p3Precio := "Ratón Óptico", 12.00

	// Voy a utilizar unos parámetros del método format, que me permite posicionar bien el texto.
	// Imprimo la cabecera de la tabla.
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("%-3s | %-18s | %-8s\n", "Nº", "PRODUCTO", "PRECIO")
	fmt.Println("------------------------------------------------------------------")

	// Muestro cada fila de datos de los productos.
	fmt.Printf("%-3d	|	%-18s	|	%8.2f€\n", 1, p1Nombre, p1Precio)
	fmt.Printf("%-3d	|	%-18s	|	%8.2f€\n", 2, p2Nombre, p2Precio)
	fmt.Printf("%-3d	|	%-18s	|	%8.2f€\n", 3, p3Nombre, p3Precio)

	// Almaceno en una variable la suma del precio de los productos.
	subtotal := p1Precio + p2Precio + p3Precio

	// Aplcio el impuesto sobre el subtotal para obtener la cifra final.
	totalConIva := subtotal * (1 + IVA)

	// Imprimo el resultado final, aplicando el iva.
	fmt.Println("------------------------------------------------")
	fmt.Printf("Subtotal: %28.2f€\n", subtotal)
	fmt.Printf("Total (con IVA): → %17.2f€\n", totalConIva)
	fmt.Println("------------------------------------------------")
}
