package main

import (
	"fmt"
	"time"
)

var dayOrder int

func main() {

	printWeekdays(loadWeekDays())

	weekDay := loadWeekDays()

	requestWeekDay()
	printRequestedDayFrom(weekDay)

}

func requestWeekDay() {
	fmt.Print("Insert day order: ")
	fmt.Scanln(&dayOrder)
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

func printRequestedDayFrom(w map[int]string) {
	println(w[dayOrder])
}
