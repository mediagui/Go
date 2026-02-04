package main

import (
	"fmt"
	"time"
)

// Example documentation
// for main func
func main() {

	printWeekdays(getWeekDaysMap())

	weekDayMap := getWeekDaysMap()
	requestedDay := requestWeekDayFromConsole()

	printRequestedDayFrom(weekDayMap, requestedDay)

}

func requestWeekDayFromConsole() int {

	var dayOrder int

	fmt.Print("Insert day order: ")
	fmt.Scanln(&dayOrder)
	return dayOrder
}

/*
Function comments
another function comment
*/
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

func printRequestedDayFrom(w map[int]string, dayOrder int) {
	println(w[dayOrder])
}
