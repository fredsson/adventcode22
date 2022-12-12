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
	div            int
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

func OperationFromInput(input string, feelRelief bool) func(int) int {
	if strings.Contains(input, "*") {
		rightHand := strings.Split(input, "*")[1]
		if strings.Contains(rightHand, "old") {
			return func(i int) int {
				result := i * i
				if feelRelief {
					return int(math.Floor(float64(result) / 3.))
				}
				return result
			}
		} else {
			v, _ := strconv.Atoi(strings.TrimSpace(rightHand))
			return func(i int) int {
				result := i * v
				if feelRelief {
					return int(math.Floor(float64(result) / 3.))
				}
				return result
			}
		}
	}
	if strings.Contains(input, "+") {
		rightHand := strings.Split(input, "+")[1]
		v, _ := strconv.Atoi(strings.TrimSpace(rightHand))
		return func(i int) int {
			result := i + v
			if feelRelief {
				return int(math.Floor(float64(result) / 3.))
			}
			return result
		}
	}

	return func(i int) int {
		return i
	}
}

func TestFromInput(input string) (int, func(int) bool) {
	parts := strings.Split(input, " ")
	div, _ := strconv.Atoi(parts[len(parts)-1])

	return div, func(i int) bool {
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

func CreateMonkey(lines []string, feelRelief bool) *Monkey {
	monkey := new(Monkey)

	monkey.items = StartingItemsFromInput(lines[1])
	monkey.op = OperationFromInput(lines[2], feelRelief)

	monkey.div, monkey.test = TestFromInput(lines[3])

	monkey.buddies = BuddiesFromInput(lines[4], lines[5])
	monkey.inspectedItems = 0

	return monkey
}

func (monkey *Monkey) PlayTurn(monkeys []*Monkey) {
	for _, item := range monkey.items {
		monkey.inspectedItems++
		newItem := monkey.op(item)
		if monkey.test(newItem) {
			monkeys[monkey.buddies[0]].items = append(monkeys[monkey.buddies[0]].items, newItem)
		} else {
			monkeys[monkey.buddies[1]].items = append(monkeys[monkey.buddies[1]].items, newItem)
		}
	}
	monkey.items = []int{}
}

func (monkey *Monkey) PlayTurnB(monkeys []*Monkey, lcd int) {
	for _, item := range monkey.items {
		monkey.inspectedItems++
		newItem := monkey.op(item) % lcd
		if monkey.test(newItem) {
			monkeys[monkey.buddies[0]].items = append(monkeys[monkey.buddies[0]].items, newItem)
		} else {
			monkeys[monkey.buddies[1]].items = append(monkeys[monkey.buddies[1]].items, newItem)
		}
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
	monkeysB := []*Monkey{}
	lines := []string{}
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()
		if len(lines) > 0 && strings.Contains(input, "Monkey") {
			monkeys = append(monkeys, CreateMonkey(lines, true))
			monkeysB = append(monkeysB, CreateMonkey(lines, false))
			lines = []string{}
		}

		lines = append(lines, input)
	}
	monkeys = append(monkeys, CreateMonkey(lines, true))
	monkeysB = append(monkeysB, CreateMonkey(lines, false))

	for round := 0; round < 20; round++ {
		for _, monkey := range monkeys {
			monkey.PlayTurn(monkeys)
		}
	}

	lcd := 1
	for _, monkey := range monkeysB {
		lcd *= monkey.div
	}

	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeysB {
			monkey.PlayTurnB(monkeysB, lcd)
		}
	}
	for _, monkey := range monkeysB {
		fmt.Println(monkey.inspectedItems)
	}

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItems < monkeys[j].inspectedItems
	})

	sort.SliceStable(monkeysB, func(i, j int) bool {
		return monkeysB[i].inspectedItems < monkeysB[j].inspectedItems
	})

	mostActive := monkeys[len(monkeys)-1]
	secondMostActive := monkeys[len(monkeys)-2]
	monkeyBusiness := mostActive.inspectedItems * secondMostActive.inspectedItems

	mostActiveB := monkeysB[len(monkeysB)-1]
	secondMostActiveB := monkeysB[len(monkeysB)-2]
	monkeyBusinessB := mostActiveB.inspectedItems * secondMostActiveB.inspectedItems

	return monkeyBusiness, monkeyBusinessB
}
