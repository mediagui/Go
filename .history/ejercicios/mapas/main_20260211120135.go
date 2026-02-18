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

		newDay := askIntInRange("Indica el nuevo día por su índice [0-6]: ", 0, 6)

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

	weekDays[int(dayToBeChanged)] = time.Weekday(newDay).String()
}

// requestUserTheDayToChange muestra al usuario el mapa actual y solicita
// el índice del día que desea cambiar. Devuelve el valor como `dayToChange`.
func requestUserTheDayToChange() dayToChange {
	weekDays := fmt.Sprintf("%v", getWeekDays())
	idx := askIntInRange("Elije el día a cambiar en el mapa "+weekDays+" (0-6): ", 0, 6)
	return dayToChange(idx)
}

// fillTheMapWithWeekDays rellena el mapa `weekDays` con los nombres de
// los días de la semana usando índices 1..7 (1=Lunes ... 7=Domingo).
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
	fmt.Println(question)
	fmt.Scanln(&answer)
	return answer
}

// askIntInRange solicita al usuario un entero y repite hasta que esté
// dentro del rango [min,max]. Devuelve el entero validado.
func askIntInRange(question string, min int, max int) int {
	for {
		var v int
		fmt.Println(question)
		_, err := fmt.Scanln(&v)
		if err != nil {
			fmt.Printf("Entrada inválida. Introduce un número entre %d y %d.\n", min, max)
			continue
		}
		if v < min || v > max {
			fmt.Printf("Valor fuera de rango. Debe estar entre %d y %d.\n", min, max)
			continue
		}
		return v
	}
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
