package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type OpenFile struct {
	File    *os.File
	Scanner *bufio.Scanner
}

func readFileByLines(filename string) *OpenFile {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	openFile := new(OpenFile)
	openFile.File = readFile
	openFile.Scanner = fileScanner

	return openFile
}

func DayOne() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d1.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	var caloriesByElf []int
	currentCalories := 0
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()
		if input == "" {
			caloriesByElf = append(caloriesByElf, currentCalories)
			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Could not convert %s to calories", input)
			continue
		}

		currentCalories += calories
	}

	sort.Ints(caloriesByElf)

	a := caloriesByElf[len(caloriesByElf)-1]

	b := caloriesByElf[len(caloriesByElf)-1] + caloriesByElf[len(caloriesByElf)-2] + caloriesByElf[len(caloriesByElf)-3]

	openFile.File.Close()
	return a, b
}
