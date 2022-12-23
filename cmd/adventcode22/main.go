package main

import (
	"fmt"
	"time"
)

func main() {
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
		DayTwenty,
		DayTwentyone,
	}

	total := time.Now()
	for index, dayFunc := range days {
		start := time.Now()
		a, b := dayFunc()
		duration := time.Since(start)
		fmt.Printf("------ day %d (%s) -----\n", index+1, duration.Truncate(time.Microsecond))
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println("-------")
	}
	fmt.Println("-------")
	fmt.Println("Total:", time.Since(total))
}
