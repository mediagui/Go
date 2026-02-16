// Paquete del proyecto.
package manejoErrores

import (
	"fmt"
)

/*
Al poner el defer entodos, se ejecutaría de la siguiente manera

defer fmt.Println(2)
defer fmt.Println(1)
defer fmt.Println(3)
*/

func ExplDefer() {
	/*
		La plabra clave [defer], se utiliza para posponer la ejecución de una función hasta que al función que lo contiene haya finalizado.
	*/
	defer fmt.Println(3)
	defer fmt.Println(1)
	defer fmt.Println(2)
}
