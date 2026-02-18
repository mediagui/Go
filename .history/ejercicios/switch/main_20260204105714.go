package main

import "time"

func main() {

	weekDay := loadWeekDays()

	println(time.Weekday.String(time.Friday))

	println(time.Monday.String())

}

func loadWeekDays() map[int]string {
	weekDay := map[int]string{}

	for i := range 7 {

		weekDay
		weekDay = append(weekDay, time.Weekday(i).String())
	}

	return weekDay
}
