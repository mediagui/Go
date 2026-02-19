// Paquete al que pertenece
package agradecimientos

// Importaciones.
import (
	"regexp"
	"testing"
)

/*
Para crear funciones de prueba, el nombre de la función debe empezar por la palabra [Test] (TestNombreFuncion). Cuando creamos las funciones de prueba, estas reciben un objeto de tipo "testing". Por lo que en este caso, tendrémos como parámetro una variable [t], que almacenará un puntero a [*testing.T], este objeto nos sirve para reportar el resultado de la prueba.
*/

// Función de prueba para la función [Hola()], que recibe un nombre y devuelve un saludo.
func TestHolaNombre(t *testing.T) {
	/*
		Parámetros del paquete testing.

		[testing.T] (Test)

			Su función es verificar la lógica y corrección del código. Si algo no da el resultado esperado, marca el test como "fallido"

		[testing.B] (Benchmark)

			Se utiliza para medir el rendimiento. Ejecuta el código muchas veces para decirnos cuántos nanosegundos tarda en cada operación.

		[testing.F] (Fuzz)

			"Experto en resitencia". Genera datos aleatorios y "extrañós" automáticamente para intentar encontrar errores de seguridad o casos límite que al programador se le haya pasado por alto.

		[testing.M] (Main)

			Se usa para configurar el entorno antes de que corra cualquier prueba (como conectar una base de datos) y limpiar todo al finalizar.
	*/
	// Preparo el sujeto de prueba.
	nombre := "Sergio"

	// Configuro la búsqueda con una expresión regulat. [\b] asegura que busque la palabra exacta.
	require := regexp.MustCompile(`\b` + nombre + `\b`)

	// Invoco a la función [Hola()] y capturo tanto su respuesta como cualquier posible error.
	msg, err := Hola("Sergio")

	// Analizo el resultado: si el nombre no coincide o si la función devuelve un error inesperado, hago un [Fatal] y detengo la prueba, generando un informe.
	if !require.MatchString(msg) || err != nil {
		t.Fatalf(`Hola("Sergio") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, require)
	}
	/*
		Resumen de "Verbos" de Formato en Fatalf

		[%q]

			Envuelve el texto en comillas para detectar espacios invisibles o cadenas vacías.

		[%v]

			Muestra el valor del erro de forma natural

		[%#q]

			Muestra la expresión regular tal cual está escrita en el código para depurar el patrón.
	*/
} // Si llego aquí sin errores, el visto bueno al código.

// Prueba para casos vacíos. Quiero verificar que el sistema "sepa quejarse" cuando no recibe datos.
func TestHolaVacio(t *testing.T) {
	// Creo un "cebo" vacío con la función [Hola()].
	msg, err := Hola("")
	// La logica que utilizo aquí es la inversa a la anterior. Si devuelve un mensaje (no está vacío) o si NO hay error (es nil), emito un "fallo fatal".
	if msg != "" || err == nil {
		t.Fatalf(`Hola("") = %q, %v, quiere "", error`, msg, err)
	}
} // Si la función falló como debía, considero que he terminado con éxito.

// Paquete al que pertenece
// package agradecimientos

// // Importaciones.
// import (
// 	"regexp"
// 	"testing"
// )

// /*
// Para crear funciones de prueba, el nombre de la función debe empezar por la palabra [Test] (TestNombreFuncion). Cuando creamos las funciones de prueba, estas reciben un objeto de tipo "testing". Por lo que en este caso, tendrémos como parámetro una variable [t], que almacenará un puntero a [*testing.T], este objeto nos sirve para reportar el resultado de la prueba.
// */

// // Función de prueba para la función [Hola()], que recibe un nombre y devuelve un saludo.
// func TestHolaNombre(t *testing.T) {
// 	/*
// 		Parámetros del paquete testing.

// 		[testing.T] (Test)

// 			Su función es verificar la lógica y corrección del código. Si algo no da el resultado esperado, marca el test como "fallido"

// 		[testing.B] (Benchmark)

// 			Se utiliza para medir el rendimiento. Ejecuta el código muchas veces para decirnos cuántos nanosegundos tarda en cada operación.

// 		[testing.F] (Fuzz)

// 			"Experto en resitencia". Genera datos aleatorios y "extrañós" automáticamente para intentar encontrar errores de seguridad o casos límite que al programador se le haya pasado por alto.

// 		[testing.M] (Main)

// 			Se usa para configurar el entorno antes de que corra cualquier prueba (como conectar una base de datos) y limpiar todo al finalizar.
// 	*/
// 	// Preparo el sujeto de prueba.
// 	nombre := "Sergio"

// 	// Configuro la búsqueda con una expresión regulat. [\b] asegura que busque la palabra exacta.
// 	require := regexp.MustCompile(`\b` + nombre + `\b`)

// 	// SABOTAJE 1: Le pido un saludo para "Carlos", cuando el espera "Sergio".
// 	// Invoco a la función [Hola()] y capturo tanto su respuesta como cualquier posible error.
// 	msg, err := Hola("Carlos")

// 	// Analizo el resultado: si el nombre no coincide o si la función devuelve un error inesperado, hago un [Fatal] y detengo la prueba, generando un informe.
// 	if !require.MatchString(msg) || err != nil {
// 		t.Fatalf(`Hola("Sergio") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, require)
// 	}
// 	/*
// 		Resumen de "Verbos" de Formato en Fatalf

// 		[%q]

// 			Envuelve el texto en comillas para detectar espacios invisibles o cadenas vacías.

// 		[%v]

// 			Muestra el valor del erro de forma natural

// 		[%#q]

// 			Muestra la expresión regular tal cual está escrita en el código para depurar el patrón.
// 	*/
// } // Si llego aquí sin errores, el visto bueno al código.

// // Prueba para casos vacíos. Quiero verificar que el sistema "sepa quejarse" cuando no recibe datos.
// func TestHolaVacio(t *testing.T) {
// 	// Sabotaje 2: Le envío un nombre válido ("Sergio") a la prueba que debería estar vacía.
// 	// Creo un "cebo" vacío con la función [Hola()].
// 	msg, err := Hola("Sergio") // Devolverá un mensaje y NO  devolverá error (err será [nil])
// 	// La logica que utilizo aquí es la inversa a la anterior. Si devuelve un mensaje (no está vacío) o si NO hay error (es nil), emito un "fallo fatal".
// 	if msg != "" || err == nil {
// 		t.Fatalf(`Hola("") = %q, %v, quiere "", error`, msg, err)
// 	}
// } // Si la función falló como debía, considero que he terminado con éxito.
