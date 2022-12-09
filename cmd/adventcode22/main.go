package main

import (
	"fmt"
	"time"
)

func main() {
	days := []func(){
		dayOne,
		DayTwo,
		dayThree,
		DayFour,
		DayFive,
		DaySix,
		DaySeven,
		DayEight,
		DayNine,
	}

	for index, dayFunc := range days {
		start := time.Now()
		fmt.Println("------ day", index+1, "---------")
		dayFunc()
		duration := time.Since(start)
		fmt.Println("duration:", duration)
	}
}
