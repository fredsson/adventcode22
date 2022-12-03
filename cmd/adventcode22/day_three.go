package main

import (
	"fmt"
	"strings"
)

type Rucksack struct {
	Content          string
	LeftCompartment  string
	RightCompartment string
}

var SYMBOL_BY_PRIORITY = map[int]string{
	1:  "a",
	2:  "b",
	3:  "c",
	4:  "d",
	5:  "e",
	6:  "f",
	7:  "g",
	8:  "h",
	9:  "i",
	10: "j",
	11: "k",
	12: "l",
	13: "m",
	14: "n",
	15: "o",
	16: "p",
	17: "q",
	18: "r",
	19: "s",
	20: "t",
	21: "u",
	22: "v",
	23: "w",
	24: "x",
	25: "y",
	26: "z",

	27: "A",
	28: "B",
	29: "C",
	30: "D",
	31: "E",
	32: "F",
	33: "G",
	34: "H",
	35: "I",
	36: "J",
	37: "K",
	38: "L",
	39: "M",
	40: "N",
	41: "O",
	42: "P",
	43: "Q",
	44: "R",
	45: "S",
	46: "T",
	47: "U",
	48: "V",
	49: "W",
	50: "X",
	51: "Y",
	52: "Z",
}

func CreateRucksack(content string) *Rucksack {
	rucksack := new(Rucksack)
	rucksack.Content = content

	middle := len(content) / 2
	rucksack.LeftCompartment = content[0:middle]
	rucksack.RightCompartment = content[middle:]

	return rucksack
}

func (rucksack *Rucksack) PriorityOfMissplacedItems() int {
	totalPriority := 0
	for priority := 1; priority <= len(SYMBOL_BY_PRIORITY); priority++ {
		symbol, found := SYMBOL_BY_PRIORITY[priority]
		if found && strings.Contains(rucksack.LeftCompartment, symbol) && strings.Contains(rucksack.RightCompartment, symbol) {
			totalPriority += priority
		}
	}
	return totalPriority
}

func (rucksack *Rucksack) findPriorityForMatchingItem(second *Rucksack, third *Rucksack) int {
	for priority := 1; priority <= len(SYMBOL_BY_PRIORITY); priority++ {
		symbol, found := SYMBOL_BY_PRIORITY[priority]
		if found && strings.Contains(rucksack.Content, symbol) && strings.Contains(second.Content, symbol) && strings.Contains(third.Content, symbol) {
			return priority
		}
	}
	return 0
}

func dayThree() {
	openFile := readFileByLines("inputs/d3.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return
	}

	totalPriority := 0
	var group [3]*Rucksack
	groupIndex := 0
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()
		group[groupIndex] = CreateRucksack(input)
		groupIndex++
		if groupIndex == 3 {
			totalPriority += group[0].findPriorityForMatchingItem(group[1], group[2])
			groupIndex = 0
		}

		/*priority := CreateRucksack(input).PriorityOfMissplacedItems()

		totalPriority += priority*/
	}

	fmt.Println(totalPriority)
}
