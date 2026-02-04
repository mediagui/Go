package main

import (
	"fmt"
	"time"
)

func main() {

	printWeekdays(getWeekDaysMap())

	weekDayMap := getWeekDaysMap()
	requestedDay := requestWeekDayFromConsole()

	printRequestedDayFrom(requestedDay, weekDayMap)

}

func requestWeekDayFromConsole() int {

	var dayOrder int

	fmt.Print("Insert day order: ")
	fmt.Scanln(&dayOrder)
	return dayOrder
}

func getWeekDaysMap() map[int]string {
	weekDay := map[int]string{}

	for i := range 7 {

		weekDay[i] = time.Weekday(i).String()

	}

	return weekDay
}

func printWeekdays(w map[int]string) {

	for v, k := range w {

		println(v, k)

	}

}

func printRequestedDayFrom(dayOrder int, w map[int]string) {
	println(w[dayOrder])
}
