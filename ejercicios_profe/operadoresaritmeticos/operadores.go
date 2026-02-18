// Indico el paquete donde me encuentro.
package operadoresaritmeticos

// Importaciones.
import (
	// Importo el paquete [fmt] para el trato con la consola.
	"fmt"
	// Importo el paquete [math] para realizar operaciones matemáticas avanzadas.
	"math"
)

/*
###########
#VARIABLES#
###########
*/
// Declaro las variables [a] y [b] de tipo entero, que almacenarán valores introducidos por el usuario.

var (
	a         int
	b         int
	c         int
	base      float64
	exponente float64
)

// Función pública para tratar con los distintos operadores. Esta función no recive ni devuelve ningún valor.
func MostrarOperaciones() {
	// Imprimo el título del programa.
	fmt.Println("##############################")
	fmt.Println("#Operadores Aritméticos en Go#")
	fmt.Println("##############################\n")

	// Imprimo una tabla con los operadores aritméticos básicos disponibles en Go.
	fmt.Println("+------------------------------------+")
	fmt.Println("|     Operador      |      Nombre    |")
	fmt.Println("+------------------------------------+")
	fmt.Println("|         +         |       Suma     |")
	fmt.Println("|         -         |       Resta    |")
	fmt.Println("|         *         | Multiplicación |")
	fmt.Println("|         /         |     División   |")
	fmt.Println("|         %         |      Módulo    |")
	fmt.Println("+------------------------------------+")

	// Solicito al usuario que ingrese el primer número y lo almaceno en una variable.
	fmt.Print("Ingrese el número para la variable [a] → ")
	// Utilizo [Scanln] para leer la entrada del usuario y asignarle el valor a la variable [a]
	fmt.Scanln(&a)

	// Solicito al usuario que ingrese el segundo número y lo almaceno en una variable.
	fmt.Print("Ingrese el número para la variable [b] → ")
	// Utilizo [Scanln] para leer la entrada del usuario y asignarle el valor a la variable [b]
	fmt.Scanln(&b)

	// Imprimo los números ingresados por el usuario.
	fmt.Printf("Has ingresado los números [a = %d] y [b = %d]\n", a, b)

	// Realizo y muestro las operaciones aritméticas básicas entre [a] y [b]
	fmt.Printf("\n[Suma: a + b] → %d + %d = %d\n", a, b, a+b)
	fmt.Printf("[Resta: a - b] → %d - %d = %d\n", a, b, a-b)
	fmt.Printf("[Multiplicación: a * b] → %d * %d = %d\n", a, b, a*b)
	// En caso de la división, convierto los operadores a tipo [float64] para que se muestren los números con decimales. También tengo que indicar el modificador con [%f], para que se imprima el númer decimal. Para mostrar el número concreto de decimales deseado, lo indicaríamos de esta manera [%.nf].
	fmt.Printf("[División: a / b] → %d / %d = %.2f\n", a, b, float64(a)/float64(b))
	fmt.Printf("[Módulo: a %% b] → %d %% %d = %d\n", a, b, a%b)

	/*
		También nos podemos encontrar con los operadores de incremento (++) y de decremento (--).
		Estos operadores se utilizan para aumentar o disminuir el valor de una variable en 1, respectivamente..
		No devuelven ningún valor y solo se pueden ser usados como expresiones independientes.
	*/
	fmt.Println("\n#####################################")
	fmt.Println("#Operadores de Incremento/Decremento#")
	fmt.Println("#####################################")
	// Imprimo una tabla con los operadores de incremento y decremento.
	fmt.Println("\n+--------------------------------------------------+")
	fmt.Println("|      Operador     |      Nombre    | Explicación |")
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("|         ++        |   Incremento   |  n = n + 1  |")
	fmt.Println("|         --        |   Decremento   |  n = n - 1  |")
	fmt.Println("+--------------------------------------------------+")

	// Muestro el valor original de [b] y luego aplico el operador de incremento y muestro el nuevo valor.
	fmt.Printf("\nValor original de [b]: %d\n", b)
	b++
	fmt.Printf("Valor de [b] después de [b++]: %d\n", b)

	// Muestro el valor original de [a] y luego aplico el operador de incremento y muestro el nuevo valor.
	fmt.Printf("\nValor original de [a]: %d\n", a)
	a--
	fmt.Printf("Valor de [a] después de [a--]: %d\n", a)

	/*
		Otro tipo de operador, son los operadores de asignación. Estos operadores combinan una operación aritmética con la asignación de un valor a una variable.

		Los operadores de asignación son:

		=
		+=
		-=
		*=
		/=
		%=
	*/

	fmt.Println("\n##########################")
	fmt.Println("#Operadores de Asignación#")
	fmt.Println("##########################")

	// Imprimo una tabla con los operadores de asignación.
	fmt.Println("\n+------------------------------------+")
	fmt.Println("|      Operador     |      Nombre    |")
	fmt.Println("+------------------------------------+")
	fmt.Println("|         =         |   Asignación   |")
	fmt.Println("|        +=         | Suma y Asign.  |")
	fmt.Println("|        -=         | Resta y Asign. |")
	fmt.Println("|        *=         | Mult. y Asign. |")
	fmt.Println("|        /=         | Div. y Asign.  |")
	fmt.Println("|        %=         | Módulo y Asign.|")
	fmt.Println("+------------------------------------+")

	// Pido al usuario un valor para la variable [c] y lo almaceno.
	fmt.Print("\nIngrese un valor para [c] → ")
	fmt.Scanln(&c)
	// Muestro el valor de la variable [c].
	fmt.Printf("Usted ha ingresado el valor [%d]\n", c)
	// Asigno a la variable [c], el resultado de sumarle 10 a su valor original.
	c += 10
	fmt.Printf("El valor de [c] después de [c += 10]: %d (c = c +10)\n", c)
	// Asigno a la variable [c], el resultado de restale 5.
	c -= 5
	fmt.Printf("El valor de [c] después de [c -= 5]: %d (c = c - 5)\n", c)
	// Asigno a la variable [c], el resultado de multiplicarle 2.
	c *= 2
	fmt.Printf("El valor de [c] después de [c *= 2]: %d (c = c * 2)\n", c)
	// Asigno a la variable [c], el resultado de dividirle 3.
	c /= 3
	fmt.Printf("El valor de [c] después de [c /= 2]: %d (c = c / 3)\n", c)
	// Asigno a la variable [c], el resultado del módulo del valor actual entre 2.
	c %= 3
	fmt.Printf("El valor de [c] después de [c %%= 2]: %d (c = c %% 3)\n", c)

	/*
		Con esto, ya podemos resolver operaciones matemáticas básicas, pero para operaciones más complejas, podemos utilizar el paquete "math" que nos ofrece funciones y constantes matemáticas avanzadas.

		Si deseas profundizar más en los operadores y sus usos, te invito a explorar la documentación oficial de Go en: https://golang.org/

		Este paquete nos permite trabajar co operaciones más compleas y realizar cálculos matemáticos avanzados. También nos brinda constantes matemáticas útiles como Pi, e, entre otras.
	*/

	fmt.Println("\n################")
	fmt.Println("#Paquete [math]#")
	fmt.Println("################")

	// Utilizo el paquete [math] para mostrar el valor de la constante [Pi]
	fmt.Printf("Valor de [Pi]: %f\n", math.Pi)

	/*
		También podemos imprimir los rangos de los valores numéricos que podemos utilizar en Go.
	*/
	fmt.Println("\n####################################")
	fmt.Println("#Rangos de Valores Numéricos en Go#")
	fmt.Println("####################################")
	// Imprimo los rangos de los tipos de datos numéricos en Go.
	fmt.Println("\n+-----------------------------------------------------------------------+")
	fmt.Println("|      Tipo de Dato      | Rango                    ")
	fmt.Println("+-----------------------------------------------------------------------+")
	fmt.Printf("|        int8            → [%d  a  %d]\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("|        int16           → [%d  a  %d]\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("|        int32           → [%d  a  %d]\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("|        int64           → [%d  a  %d]\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("|        uint8           → [%d  a  %d]\n", 0, math.MaxUint8)
	fmt.Printf("|        uint16          → [%d  a  %d]\n", 0, math.MaxUint16)
	fmt.Printf("|        uint32          → [%d  a  %d]\n", 0, math.MaxUint32)
	fmt.Printf("|        uint64          → [0  a  %.0e]\n", float64(math.MaxUint64))
	fmt.Println("+-----------------------------------------------------------------------+")

	/*
		También podemos conseguir potencias y raíces utilizano la funciones de [math]
	*/

	fmt.Println("\n#########################################")
	fmt.Println("#Operaciones de Potencia y Raíz Cuadrada#")
	fmt.Println("#########################################")
	// Solicito al usuario que ingrese la base y el exponente para calcular la potencia.
	fmt.Print("\nIngrese la base para la potencia → ")
	fmt.Scanln(&base)
	fmt.Print("\nIngrese el exponente para la potencia → ")
	fmt.Scanln(&exponente)
	// Calculo la potencia usando la función [Pow()] del paquete [math].
	resultado := math.Pow(base, exponente)
	// Imprimo el resultado de la potencia.
	fmt.Printf("\nEl resultado de %.2f elevado a %.2f es: %.2f\n", base, exponente, resultado)
	// Solicito al usuario que ingrese un número para calcular su raíz cuadrada.
	var numero float64
	fmt.Print("\nIngrese un número para calcular su raíz cuadrada:\n→ ")
	fmt.Scanln(&numero)
	// Calculo la raíz cuadrada usando la función Sqrt del paquete math.
	raiz := math.Sqrt(numero)
	fmt.Printf("La raíz cuadrada de %.2f es: %.2f\n", numero, raiz)
}
