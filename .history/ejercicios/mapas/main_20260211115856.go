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
	"time"
)

// dayToChange es un alias para facilitar la lectura cuando se solicita
// al usuario el índice del día a modificar.
type dayToChange int

// main es el punto de entrada del programa. Crea un mapa con los días de
// la semana, pregunta al usuario cuántos días quiere cambiar y procesa
// cada modificación antes de imprimir el mapa resultante.
func main() {
	var days map[int]string = make(map[int]string)

	fillTheMapWithWeekDays(days)

	howManyDays := makeAQuestionAndGetTheAnswer[int]("¿Cuántos días quieres cambiar: ")

	for i := 0; i < howManyDays; i++ {
		dayToBeChanged := requestUserTheDayToChange()

		newDay := makeAQuestionAndGetTheAnswer[int]("Indica el nuevo día por su índice [1-7]: ")

		setTheNewWeekDayIndex(days, dayToBeChanged, newDay)
	}

	printMap(days)
}

// setTheNewWeekDayIndex actualiza el nombre del día en el mapa `weekDays`
// en la posición indicada por `dayToBeChanged`. `newDay` espera un índice
// en el rango 1..7; se normaliza usando módulo para mapear 7 -> Domingo.
func setTheNewWeekDayIndex(weekDays map[int]string, dayToBeChanged dayToChange, newDay int) {
	weekDays[int(dayToBeChanged)] = time.Weekday(newDay % 7).String()
}

// requestUserTheDayToChange muestra al usuario el mapa actual y solicita
// el índice del día que desea cambiar. Devuelve el valor como `dayToChange`.
func requestUserTheDayToChange() dayToChange {
	weekDays := fmt.Sprintf("%v", getWeekDays())
	return makeAQuestionAndGetTheAnswer[dayToChange]("Elije el día a cambiar en el mapa " + weekDays + ": ")
}

// fillTheMapWithWeekDays rellena el mapa `weekDays` con los nombres de
// los días de la semana usando índices 1..7 (1=Lunes ... 7=Domingo).
func fillTheMapWithWeekDays(weekDays map[int]string) {
	for i := 1; i <= 7; i++ {
		weekDays[i] = time.Weekday(i % 7).String()
	}
}

// makeAQuestionAndGetTheAnswer muestra `question` por consola y lee la
// respuesta del usuario, devolviéndola con el tipo genérico `T`.
func makeAQuestionAndGetTheAnswer[T any](question string) T {
	var answer T
	fmt.Println(question)
	fmt.Scanln(&answer)
	return answer
}

// printMap imprime el mapa de días en un formato sencillo.
func printMap(weekDaysMap map[int]string) {
	fmt.Printf("Nuevo mapa: %v.\n", weekDaysMap)
}

// getWeekDays devuelve un mapa nuevo con los días de la semana (índices 1..7).
func getWeekDays() map[int]string {
	var weekDays map[int]string = make(map[int]string)

	for i := 1; i <= 7; i++ {
		weekDays[i] = time.Weekday(i % 7).String()
	}

	return weekDays
}
