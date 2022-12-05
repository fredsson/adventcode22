package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SupplyArea struct {
	areas [][]string
}

func CreateSupplyArea(input []string) *SupplyArea {
	supplyArea := new(SupplyArea)
	supplyArea.areas = [][]string{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}
	for i := len(input) - 2; i >= 0; i-- {
		for area := 0; area < 9; area++ {
			position := (area * 4) + 1
			if input[i][position] != 32 {
				supplyArea.areas[area] = append(supplyArea.areas[area], string(input[i][position]))

			}
		}
	}
	return supplyArea
}

func (supplyArea *SupplyArea) MoveSupply(amount int, fromIndex int, toIndex int) {
	adjustedFromIndex := fromIndex - 1
	adjustedToIndex := toIndex - 1
	for i := 0; i < amount; i++ {
		indexToGrab := len(supplyArea.areas[adjustedFromIndex]) - 1
		supply := supplyArea.areas[adjustedFromIndex][indexToGrab]
		supplyArea.areas[adjustedFromIndex] = supplyArea.areas[adjustedFromIndex][:indexToGrab]
		supplyArea.areas[adjustedToIndex] = append(supplyArea.areas[adjustedToIndex], supply)
	}
}

func (supplyArea *SupplyArea) MultiMoveSupply(amount int, fromIndex int, toIndex int) {
	adjustedFromIndex := fromIndex - 1
	adjustedToIndex := toIndex - 1

	startGrabIndex := len(supplyArea.areas[adjustedFromIndex]) - amount

	supplies := supplyArea.areas[adjustedFromIndex][startGrabIndex:]
	supplyArea.areas[adjustedFromIndex] = supplyArea.areas[adjustedFromIndex][:startGrabIndex]
	supplyArea.areas[adjustedToIndex] = append(supplyArea.areas[adjustedToIndex], supplies...)
}

func DayFive() {
	openFile := readFileByLines("inputs/d5.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return
	}

	supplyAreaInput := []string{}
	moves := []string{}
	blahDone := false
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()
		if input == "" {
			blahDone = true
			continue
		}
		if !blahDone {
			supplyAreaInput = append(supplyAreaInput, input)
		} else {
			moves = append(moves, input)
		}
	}

	supplyArea := CreateSupplyArea(supplyAreaInput)
	supplyAreaB := CreateSupplyArea(supplyAreaInput)

	for _, move := range moves {
		instructions := strings.Split(move, " ")

		amountInstruction, _ := strconv.Atoi(instructions[1])
		fromInstruction, _ := strconv.Atoi(instructions[3])
		toInstruction, _ := strconv.Atoi(instructions[5])

		supplyArea.MoveSupply(amountInstruction, fromInstruction, toInstruction)

		supplyAreaB.MultiMoveSupply(amountInstruction, fromInstruction, toInstruction)
	}

	result := ""
	for _, area := range supplyArea.areas {
		top := len(area) - 1
		result += area[top]
	}

	resultB := ""
	for _, area := range supplyAreaB.areas {
		top := len(area) - 1
		resultB += area[top]
	}

	fmt.Println(result)
	fmt.Println(resultB)

	openFile.File.Close()
}
