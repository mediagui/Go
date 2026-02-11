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

type dayToChange int

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

func setTheNewWeekDayIndex(weekDays map[int]string, dayToBeChanged dayToChange, newDay int) {

	weekDays[int(dayToBeChanged)] = time.Weekday(newDay).String()

}

func requestUserTheDayToChange() dayToChange {

	weekDays := fmt.Sprintf("%v", getWeekDays())
	return makeAQuestionAndGetTheAnswer[dayToChange]("Elije el día a cambiar en el mapa " + weekDays + ": ")

}

// Fill the map with the weekdays
func fillTheMapWithWeekDays(weekDays map[int]string) {

	for i := 1; i < 6; i++ {
		weekDays[i] = time.Weekday(i + 1).String()
	}

}

func makeAQuestionAndGetTheAnswer[T any](question string) T {

	var answer T
	fmt.Println(question)
	fmt.Scanln(&answer)
	return answer

}

func printMap(weekDaysMap map[int]string) {
	fmt.Printf("Nuevo mapa: %v.\n", weekDaysMap)
}

func getWeekDays() map[int]string {

	var weekDays map[int]string = make(map[int]string)

	for i := 0; i < 7; i++ {
		weekDays[i] = time.Weekday(i).String()
	}

	return weekDays
}
