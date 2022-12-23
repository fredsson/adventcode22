package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	var dayToRun int = -1
	if len(os.Args) > 1 {
		input, err := strconv.Atoi(os.Args[1])
		if err == nil {
			dayToRun = input
		}
	}

	days := []func() (interface{}, interface{}){
		DayOne,
		DayTwo,
		dayThree,
		DayFour,
		DayFive,
		DaySix,
		DaySeven,
		DayEight,
		DayNine,
		DayTen,
		DayEleven,
		DayTwelve,
		DayThirteen,
		DayFourteen,
		DayFifteen,
		DaySixteen,
		DaySeventeen,
		DayEighteen,
		DayNineteen,
		DayTwenty,
		DayTwentyone,
	}

	dayIndex := dayToRun - 1
	if dayIndex >= 0 && dayIndex < len(days) {
		f := days[dayIndex]
		RunOneDay(dayToRun, f)
	} else {
		RunAllDays(days)
	}
}

func RunOneDay(index int, dayFunc func() (interface{}, interface{})) {
	start := time.Now()
	a, b := dayFunc()
	duration := time.Since(start)
	fmt.Printf("------ day %d (%s) -----\n", index, duration.Truncate(time.Microsecond))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("-------")
}

func RunAllDays(days []func() (interface{}, interface{})) {
	total := time.Now()
	for index, dayFunc := range days {
		RunOneDay(index+1, dayFunc)
	}
	fmt.Println("-------")
	fmt.Println("Total:", time.Since(total))
}
