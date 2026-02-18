// Crear un mapa que tenga como claves los números del día de la semana y como valores el nombre del día de la semana
// correspondiente:
// 1 → Lunes
// 2 → Martes
// 3 → Miércoles
// 4 → Jueves
// 5 → Viernes
// 6 → Sábado
// 7 → Domingo
// Luego, crear un bucle que pregunte al usuario si quiere modificar algún día de la semana y cuántos. Si es así,
// le preguntará el número del día y el nuevo nombre del día. Finalmente, imprimirá el diccionario modificado.

package main

import (
	"fmt"
	"strings"
	"time"
)

// Definimos códigos de color ANSI
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

// dayToChange es un alias para facilitar la lectura cuando se solicita
// al usuario el índice del día a modificar.
type dayToChange int

// main es el punto de entrada del programa. Crea un mapa con los días de
// la semana, pregunta al usuario cuántos días quiere cambiar y procesa
// cada modificación antes de imprimir el mapa resultante.
func main() {
	var days map[int]string = make(map[int]string)

	// Añadimos valores al mapa
	fillTheMapWithWeekDays(days)

	// Preguntamos cuantos días queremos cambiar
	howManyDays := makeAQuestionAndGetTheAnswer[int]("¿Cuántos días quieres cambiar (0-7)? ")
	// howManyDays := makeQuestion("¿Cuántos días quieres cambiar (0-7)? ", 0, 7)

	for range howManyDays {

		// Pedimos el día que queremos reemplazar
		dayToBeChanged := requestUserTheDayToChange()

		// Pedimos el nuevo día (Valor que sustituirá al anterior en el mapa)
		newDay := makeAQuestionAndGetTheAnswer[int]("Indica el nuevo día por su índice [0-6]: ")
		// newDay := makeQuestion("Indica el nuevo día por su índice [0-6]: ", 0, 6)

		// Hacemos el reemplazo activo
		setTheNewWeekDayIndex(days, dayToBeChanged, newDay)
	}

	printMap(days)
}

// setTheNewWeekDayIndex actualiza el nombre del día en el mapa `weekDays`
// en la posición indicada por `dayToBeChanged`. `newDay` debe estar
// entre 0 y 6 (0=Sunday .. 6=Saturday).
func setTheNewWeekDayIndex(weekDays map[int]string, dayToBeChanged dayToChange, newDay int) {
	if int(dayToBeChanged) < 0 || int(dayToBeChanged) > 6 {
		fmt.Printf("Índice de día a cambiar inválido: %d (debe estar entre 0 y 6)\n", dayToBeChanged)
		return
	}
	if newDay < 0 || newDay > 6 {
		fmt.Printf("Nuevo índice de día inválido: %d (debe estar entre 0 y 6)\n", newDay)
		return
	}

	weekDays[int(dayToBeChanged)] = yellow + time.Weekday(newDay).String() + reset
}

// requestUserTheDayToChange muestra al usuario el mapa actual y solicita
// el índice del día que desea cambiar. Devuelve el valor como `dayToChange`.
func requestUserTheDayToChange() dayToChange {
	// Mostrar los días numerados del 1 al 7 al usuario para que sea
	// más natural; internamente usamos índices 0..6.
	weekDaysDisplay := getWeekDaysDisplay()
	idx := makeAQuestionAndGetTheAnswer[int]("Elije el día a cambiar en el mapa " + weekDaysDisplay + " (1-7): ")
	// idx := makeQuestion("Elije el día a cambiar en el mapa "+weekDaysDisplay+" (1-7): ", 1, 7)
	return dayToChange(idx - 1)
}

// getWeekDaysDisplay devuelve una representación legible de los días
// con índices 1..7 (1=Lunes ... 7=Domingo) para mostrar al usuario.
func getWeekDaysDisplay() string {
	var s strings.Builder
	for i := 1; i <= 7; i++ {
		if i > 1 {
			s.WriteString(", ")
		}
		// i%7 convierte 7 -> 0 (Sunday)
		fmt.Fprintf(&s, "%d:%s", i, time.Weekday(i%7).String())
	}
	return s.String()
}

// fillTheMapWithWeekDays rellena el mapa `weekDays` con los nombres de
// los días de la semana usando índices 1..7 (1=Lunes ... 7=Domingo).
//
// # Esto es una cabecera
//
// [Esto es un link]: https://example.org
func fillTheMapWithWeekDays(weekDays map[int]string) {
	// Usamos índices 0..6 para los días (0=Sunday .. 6=Saturday)
	for i := 0; i <= 6; i++ {
		weekDays[i] = time.Weekday(i).String()
	}
}

// makeAQuestionAndGetTheAnswer muestra `question` por consola y lee la
// respuesta del usuario, devolviéndola con el tipo genérico `T`.
func makeAQuestionAndGetTheAnswer[T any](question string) T {
	var answer T
	fmt.Print(question)
	fmt.Scanln(&answer)
	return answer
}

// makeQuestion solicita al usuario un entero y repite hasta que esté
// dentro del rango [min,max]. Devuelve el entero validado.
func makeQuestion[T any](question string, min T, max T) T {
	for {
		var v T
		fmt.Println(question)
		_, err := fmt.Scanln(&v)

		//Valida si ha habido algún problema al leer el valor desde teclado
		if err != nil {
			fmt.Printf(red+"Entrada inválida. Introduce un número entre %v y %v.\n"+reset, min, max)
			continue
		}

		//Valida que el valor se encuentre entre los límites establecidos
		//FIXME: Corregir para que funcione con genéricos
		// if v < min || v > max {
		// 	fmt.Printf("Valor fuera de rango. Debe estar entre %d y %d.\n", min, max)
		// 	continue
		// }
		return v
	}
}

// printMap imprime el mapa de días en un formato sencillo.
func printMap(weekDaysMap map[int]string) {
	// Construir la salida en una sola cadena para evitar múltiples llamadas a fmt.
	var s strings.Builder
	s.WriteString("Nuevo mapa: map[")
	for i := 0; i <= 6; i++ {
		if i > 0 {
			s.WriteString(" ")
		}
		// i+1 es 1..7 y cabe en un solo dígito -> convertir por aritmética de bytes.
		s.WriteString(string('0' + byte(i+1)))
		s.WriteString(":")
		s.WriteString(weekDaysMap[i])
	}
	s.WriteString("].")
	fmt.Println(s.String())
}

// getWeekDays devuelve un mapa nuevo con los días de la semana (índices 0..6).
func getWeekDays() map[int]string {
	var weekDays map[int]string = make(map[int]string)

	for i := 0; i <= 6; i++ {
		weekDays[i] = time.Weekday(i).String()
	}

	return weekDays
}
