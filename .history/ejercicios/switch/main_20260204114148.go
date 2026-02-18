package main

import (
	"fmt"
	"time"
)

dayOrder int

func main() {

	printWeekdays(loadWeekDays())

	weekDay := loadWeekDays()

	printRequestedDay(dayOrder)

}

func requestWeekDay(){
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
