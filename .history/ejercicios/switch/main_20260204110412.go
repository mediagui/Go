package main

import "time"

func main() {

	weekDay := loadWeekDays()

	println(time.Weekday.String(time.Friday))

	println(time.Monday.String())

	println(weekDay)

}

func loadWeekDays() map[int]string {
	weekDay := map[int]string{}

	for i := range 7 {

		weekDay[i] = time.Weekday(i).String()

	}

	return weekDay
}
