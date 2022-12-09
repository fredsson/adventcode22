package main

import "fmt"

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
		fmt.Println("------ day", index+1, "---------")
		dayFunc()
	}
}
