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

	var days map[int]string

	fillTheMapWithWeekDays(days)

	requestUserTheDayToChange()

}

func makeAQuestionAndGetTheAnswer(question string) any {

	var answer any
	fmt.Println(question)
	fmt.Scanln(&answer)
	return answer

}

func requestUserTheDayToChange() {

	howManyDays := makeAQuestionAndGetTheAnswer("¿Cuántos días quieres cambiar")

	for i:=1; i<=int(howManyDays.(int)); i++{
		makeAQuestionAndGetTheAnswer("Elije el día: ")
	}

}

func requestUserToChangeTheMap(d dayToChange, weekDays map[int]string) {

}

// Fill the map with the weekdays
func fillTheMapWithWeekDays(weekDays map[int]string) {

	for i := 1; i < 6; i++ {
		weekDays[i] = time.Weekday(i + 1).String()
	}

}
