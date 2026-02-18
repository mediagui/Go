// Declasro el paquete al que pertenece el archivo.
package agradecimientos

// Importaciones
import (
	"regexp"  // Me permite trabajar con expresiones regulares.
	"testing" // Paqute necesarioo para realizar prueba en go.
)

// Función para probar el saludo.
func TestHolaNombre(t *testing.T) {
	// Defino el nombre que espero encontrar en el saludo
	nombre := "Rigorbeto"

	// Creo una expresión regular que busca el nombnre dentro del mensaje.
	/*
		[regexp.MustCompile] compila la expresión y falla si es inválida.
		[\b] Indica el límite de la palabra, evita "falsos positivos" como "Rigorbeto123"
	*/
	quiero := regexp.MustCompile(`\b` + nombre + `\b`)

	// Llamo a la funcion Hola, pasándo el nombre almacenado en la variable.
	msg, err := Hola("")

	/*
		Aquí haré dos comprobaciones:
		1. Que el mensaje contenga el nombre usado en la expresión regular.
		2. Que no haya ocurrido ningún error.
	*/
	if !quiero.MatchString(msg) || err != nil {
		// Si algo falla, utilizo un [Fatalf] para detener la prueba y muestro un mensaje detallado.
		t.Fatalf(`Hola("Rigorbeto") = %q, %v, quiere para %#q, nil`, msg, err, quiero)
		/*
			[%q]	-> Imprime una cadena entre comillas, escapando caracteres especiales.
			[%v]	-> Imprime el valor "tal cual".
			[%#q]	-> Imprime una cadena entre comillas y con "formato Go", útil para ver exáctamente cómo está construida la expresión.
		*/
	}
}

// Función que llama a [agradecimientos.Hola] con una cadena vacía y verifica que se devuelva un error.
func TestHelloEmpty(t *testing.T) {
	// Llamo a [hola] con un string vacío.
	msg, err := Hola("")
	/*
		Verifico dos cosas:
		1. Que el mensaje devuelto esté vacío.
		2. Que se haya devuelto un error.
	*/
	if msg != "" || err == nil {
		// Ai no se cumple lo esperado, la prueba falla.
		t.Fatalf(`Hola("") = %q, %v, quiere "", error`, msg, err)
	}
}
