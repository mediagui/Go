// Defino el paquete al que pertenefe el código
package mapas

// Zona de importaciones.
import (
	"fmt"
)

/*
	Desglose de la sentencia con colores ANSI:
	1. "\033": Carácter de escape (Esc) que inicia una secuencia de control.
	2. "[38;2;...m": Configura el color del texto (foreground) en modo RGB.
	   - 38;2: Indica modo TrueColor (24-bit).
	   - 255;0;0: Valores R;G;B (Rojo puro en este caso).
	   - m: Termina la configuración del color.
	3. "..." : Texto y variables a imprimir con el color activo.
	4. "\033[0m": Secuencia de RESET. Devuelve la consola a su color original
	   para no pintar el resto de logs del programa.
	5. "\033[48;2;...m": Configura el color del fondo (background) en modo RGB.
	- 48;2: Indica modo TrueColor (24-bit).
	- 255;255;255: Valores R;G;B (Blanco puro en este caso).
	- m: Termina la configuración del color.
*/
// Zona de declaración de variables.
var (
	// Variables para almacenar los códigos de colores.
	codigoRojo   = "\033[38;2;255;0;0m"
	codigoVerde  = "\033[38;2;0;255;0m"
	codigoAzul   = "\033[38;2;0;0;255m"
	codigoBlanco = "\033[38;2;255;255;255m"
	fondoBlanco  = "\033[48;2;255;255;255m"

	// Variable para almacenar el reset de los colores.
	reset = "\033[0m"
)

// Función para explicar los mapas.
func ExplicacionMapas() {
	// Imprimo el título del programa.
	fmt.Println(codigoRojo,
		`
##################
#EJECUTANDO MAPAS#
##################
		`, reset)

	/*
		#######
		#MAPAS#
		#######

		Las mapas son como una listas que almacena cantidad de datos con clave y valor, Map en Go se parece a los diccionarios de Python.

		Para definir una mapa se hace como variables o array, en este caso ponemos map luego dentro de los corchetes ponemos el tipo de dato que será la clave y luego el tipo de dato que será su valor y podemos hacer inicializando con valor predeterminado dentro de las llaves, clave y valor separados con dos puntos y cada elemento separados con coma.

		#SINTAXIS#
		nombreMapa := map[tipoDatoClave]tipoDatoValor {
			"clave1" : valor1,
			"clave2" : valor2,
			"clave3" : valor3,
		}
	*/

	// Creo un mapa que almacenará como claves el nombre de un color y como valor su código exadecimal.
	colores := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	// Imprimo el mapa.
	fmt.Println(colores) // Como vemos, el orden que sigue a la hora de imprimir los valores, es en orden alfabético de las claves.

	// También puedo imprimir el valor de cada clave, con un bucle [for-range]
	for clave, valor := range colores {
		// Imprimo cada clave y su valor.
		// fmt.Println("La clave [" + clave + "] tiene el valor [" + valor + "]")

		// Imprimo cada clave y su valor, pintando cada uno de su color correspondiente.
		switch clave {
		case "rojo":
			fmt.Println(codigoRojo + "La clave [" + clave + "] tiene el valor [" + valor + "]" + reset)
		case "verde":
			fmt.Println(codigoVerde + "La clave [" + clave + "] tiene el valor [" + valor + "]" + reset)
		case "azul":
			fmt.Println(codigoAzul + "La clave [" + clave + "] tiene el valor [" + valor + "]" + reset)
		}
	}

	/*
		Para acceder a un único elemento del mapa, lo haremos mediante la clave.
	*/
	// Imprimo el valor de la clave "rojo"
	fmt.Println("\n" + codigoRojo + colores["rojo"] + reset)

	/*
		Para agregar un valor al mapa, se hace igual que en los arrays, indicando la clave y su valor.
	*/
	// Añado un elemento al mapa.
	colores["negro"] = "#000000"
	fmt.Println(colores)

	/*
		También puedo almacenar el valor de una clave en una variable.
	*/
	// Declaro una variable que almacenará el valor de la clave "azul"
	valor := colores["azul"]
	// Imprimo el valor de la variable.
	fmt.Printf("El valor de la clave [azul] es: %s\n", valor)

	/*
		Cuando recuperamos valores de un mapa, este devuelve dos datos, el valor y un booleano que nos indica si la clave existe o no.
	*/

	// Almaceno el valor de la clave ["rojo"] en una variable y la verificación de si existe en otra y luego imprimo ambas.
	valor, ok_rojo := colores["rojo"]
	fmt.Println(codigoRojo+valor+" : ", ok_rojo, reset)

	// Si intento recuperar un valor que no existe en el mapa, me devolverá el valor por defecto del tipo de dato de la clave y también [false]
	valor, ok_blanco := colores["blanco"]
	fmt.Println(valor+" : ", ok_blanco) // Como vemos, donde se imprime el valor, se imprime un espacio, ya que es el valor por defecto del dato de tipo [string]

	/*
		De esta forma, podemos verificar si una clave existe en el mapa.
	*/

	// Compruebo si existen ambas claves en el mapa. Utilizo un [switch] "sin etiqueta", este evelúa los casos de arriba abajo.
	switch {
	case ok_rojo && ok_blanco: // Compruebo si existen ambos valores.
		fmt.Println("Ambas claves " + codigoRojo + "[rojo]" + reset + " y [blanco] existen en el mapa")
	case ok_rojo: // Solo existe el [ok_rojo].
		fmt.Println("La clave " + codigoRojo + "[rojo]" + reset + " existe en el mapa")
	case ok_blanco: // Solo existe el [ok_blanco].
		fmt.Println(codigoVerde + "La clave " + reset + "[blanco]" + codigoVerde + " existe en el mapa" + reset)
	default: // No existen ninguno de los dos.
		fmt.Println("No existen ninguna de las dos claves en el mapa")
	}

	// También podemos realizar la verificación de si existe o no con un [if-else].
	if _, ok_verde := colores["verde"]; ok_verde {
		fmt.Println("La clave " + codigoVerde + "[verde]" + reset + " existe en el mapa.")
		fmt.Println(ok_verde)
	} else {
		// En caso de que no exista, imprimo un mensaje indicándolo.
		fmt.Println(codigoRojo + "ERROR: La clave no existe en el mapa" + reset)
		fmt.Println(ok_verde)
	}
	/*
		Para eliminar un valor del mapa, se hace mediante la clave.
	*/
	// Elimino el valor de la clave [rojo]. Pra ello, indico el mapa y la clave a eliminar: delete(mapa, clave)
	delete(colores, "rojo")
	fmt.Println(colores)

	// Imprimo nuevamente el listado de colores.
	for clave, valor := range colores {
		// Hago un [switch-case], para asegurarme de que solo imprimo cada clave y su valor, pintando cada uno de su color correspondiente.
		switch clave {
		case "rojo": // En caso de que exista la clave ["rojo"]
			fmt.Println(codigoRojo, "La clave ["+clave+"] tiene el valor ["+valor+"]"+reset) // No se imprime, ya que se eliminó.
		case "verde": // En caso de que exista la clave ["verde"]
			fmt.Println(codigoVerde, "La clave ["+clave+"] tiene el valor ["+valor+"]"+reset)
		case "azul": // En caso de que exista la clave ["azul"]
			fmt.Println(codigoAzul, "La clave ["+clave+"] tiene el valor ["+valor+"]"+reset)
		}
	}
}
