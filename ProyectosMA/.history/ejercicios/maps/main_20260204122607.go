package main

import (
	"fmt"
	"time"
)

func main() {

	printWeekdays(loadWeekDays())

	weekDayMap := loadWeekDays()
	dayOrder := requestWeekDay()
	printRequestedDayFrom(dayOrder, weekDayMap)

}

func requestWeekDay() int {

	var dayOrder int

	fmt.Print("Insert day order: ")
	fmt.Scanln(&dayOrder)
	return dayOrder
}

func loadWeekDays() map[int]string {
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
