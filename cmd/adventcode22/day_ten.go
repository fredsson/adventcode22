package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DayTen() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d10.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	cpuCommands := []string{}
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		if strings.Contains(input, "addx") {
			cpuCommands = append(cpuCommands, "noop", input)
			continue
		}

		cpuCommands = append(cpuCommands, input)
	}

	register := 1
	signal := 0
	for index, command := range cpuCommands {
		cycle := index + 1
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			signal += cycle * register
		}
		if strings.Contains(command, "addx") {
			parts := strings.Split(command, " ")
			addition, _ := strconv.Atoi(parts[1])

			register += addition
		}
	}

	openFile.File.Close()
	return signal, 0
}
