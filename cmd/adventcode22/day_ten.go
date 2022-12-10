package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CRT struct {
	row   []string
	index int
}

func CreateCRT() *CRT {
	crt := new(CRT)
	crt.row = []string{""}
	crt.index = 0

	return crt
}

func LitCycle(cycle int, register int) bool {
	if cycle == register {
		return true
	}
	if cycle == register-1 {
		return true
	}
	if cycle == register+1 {
		return true
	}
	return false
}

func (crt *CRT) SetRow(cycle int) {
	if cycle == 40 {
		crt.row = append(crt.row, "")
		crt.index = 1
	}
	if cycle == 80 {
		crt.row = append(crt.row, "")
		crt.index = 2
	}
	if cycle == 120 {
		crt.row = append(crt.row, "")
		crt.index = 3
	}
	if cycle == 160 {
		crt.row = append(crt.row, "")
		crt.index = 4
	}
	if cycle == 200 {
		crt.row = append(crt.row, "")
		crt.index = 5
	}
}

func (crt *CRT) Draw(cycle int, register int) {
	crt.SetRow(cycle)

	screen := cycle % 40

	char := "."
	if LitCycle(screen, register) {
		char = "#"
	}
	crt.row[crt.index] += char
}

func (crt *CRT) String() string {
	result := ""
	for _, row := range crt.row {
		result += row + "\n"
	}
	return result
}

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
	crt := CreateCRT()
	for index, command := range cpuCommands {
		cycle := index + 1
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			signal += cycle * register
		}
		crt.Draw(index, register)
		if strings.Contains(command, "addx") {
			parts := strings.Split(command, " ")
			addition, _ := strconv.Atoi(parts[1])

			register += addition
		}
	}

	openFile.File.Close()
	return signal, crt
}
