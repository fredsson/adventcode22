package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items          []int
	op             func(int) int
	test           func(int) bool
	buddies        [2]int
	inspectedItems int
}

func StartingItemsFromInput(input string) []int {
	split := strings.Split(input, " ")

	items := []int{}
	for _, v := range split {
		item, err := strconv.Atoi(strings.Replace(v, ",", "", 1))
		if err == nil {
			items = append(items, item)
		}
	}

	return items
}

func OperationFromInput(input string) func(int) int {
	if strings.Contains(input, "*") {
		rightHand := strings.Split(input, "*")[1]
		if strings.Contains(rightHand, "old") {
			return func(i int) int {
				return int(math.Floor(float64(i*i) / 3.))
			}
		} else {
			v, _ := strconv.Atoi(strings.TrimSpace(rightHand))
			return func(i int) int {
				return int(math.Floor(float64(i*v) / 3.))
			}
		}
	}
	if strings.Contains(input, "+") {
		rightHand := strings.Split(input, "+")[1]
		v, _ := strconv.Atoi(strings.TrimSpace(rightHand))
		return func(i int) int {
			return int(math.Floor(float64(i+v) / 3.))
		}
	}

	return func(i int) int {
		return i
	}
}

func TestFromInput(input string) func(int) bool {
	parts := strings.Split(input, " ")
	div, _ := strconv.Atoi(parts[len(parts)-1])

	return func(i int) bool {
		result := float64(i) / float64(div)
		return math.Round(result) == result
	}
}

func BuddiesFromInput(firstLine string, secondLine string) [2]int {
	result := [2]int{}

	parts := strings.Split(firstLine, " ")
	v, _ := strconv.Atoi(parts[len(parts)-1])
	result[0] = v

	parts = strings.Split(secondLine, " ")
	v, _ = strconv.Atoi(parts[len(parts)-1])
	result[1] = v

	return result
}

func CreateMonkey(lines []string) *Monkey {
	monkey := new(Monkey)

	monkey.items = StartingItemsFromInput(lines[1])
	monkey.op = OperationFromInput(lines[2])

	monkey.test = TestFromInput(lines[3])

	monkey.buddies = BuddiesFromInput(lines[4], lines[5])
	monkey.inspectedItems = 0

	return monkey
}

func (monkey *Monkey) InspectItem(item int, monkeys []*Monkey) {
	monkey.inspectedItems++
	newItem := monkey.op(item)
	if monkey.test(newItem) {
		monkeys[monkey.buddies[0]].items = append(monkeys[monkey.buddies[0]].items, newItem)
	} else {
		monkeys[monkey.buddies[1]].items = append(monkeys[monkey.buddies[1]].items, newItem)
	}
}

func (monkey *Monkey) PlayTurn(monkeys []*Monkey) {
	for _, item := range monkey.items {
		monkey.InspectItem(item, monkeys)
	}
	monkey.items = []int{}
}

func DayEleven() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d11.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	monkeys := []*Monkey{}
	lines := []string{}
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()
		if len(lines) > 0 && strings.Contains(input, "Monkey") {
			monkeys = append(monkeys, CreateMonkey(lines))
			lines = []string{}
		}

		lines = append(lines, input)
	}
	monkeys = append(monkeys, CreateMonkey(lines))

	for round := 0; round < 20; round++ {
		for _, monkey := range monkeys {
			monkey.PlayTurn(monkeys)
		}
	}

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItems < monkeys[j].inspectedItems
	})

	mostActive := monkeys[len(monkeys)-1]
	secondMostActive := monkeys[len(monkeys)-2]
	monkeyBusiness := mostActive.inspectedItems * secondMostActive.inspectedItems

	return monkeyBusiness, 0
}
