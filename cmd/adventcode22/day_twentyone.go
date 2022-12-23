package main

import (
	"fmt"
	"strconv"
	"strings"
)

type OperationFunc = func(float64, float64) float64

type CalculationMonkey struct {
	Name       string
	WaitingOn  []string
	Dependants []string
	Value      float64
	Operation  OperationFunc
}

func (monkey *CalculationMonkey) Reset() {
	monkey.WaitingOn = monkey.Dependants
	if len(monkey.Dependants) > 0 {
		monkey.Value = 0
	}
}

func NamesFromOperation(operation string, operator string) []string {
	result := []string{}
	names := strings.Split(operation, operator)
	for _, name := range names {
		result = append(result, strings.TrimSpace(name))
	}

	return result
}

func MonkeyOperationFromInput(monkey *CalculationMonkey, input string) {
	var result []string

	if strings.Contains(input, "*") {
		result = NamesFromOperation(input, "*")
		monkey.Dependants = result
		monkey.WaitingOn = result
		monkey.Operation = func(left, right float64) float64 {
			return left * right
		}
	} else if strings.Contains(input, "-") {
		result = NamesFromOperation(input, "-")
		monkey.Dependants = result
		monkey.WaitingOn = result
		monkey.Operation = func(left, right float64) float64 {
			return left - right
		}
	} else if strings.Contains(input, "+") {
		result = NamesFromOperation(input, "+")
		monkey.Dependants = result
		monkey.WaitingOn = result
		monkey.Operation = func(left, right float64) float64 {
			return left + right
		}
	} else if strings.Contains(input, "/") {
		result = NamesFromOperation(input, "/")
		monkey.Dependants = result
		monkey.WaitingOn = result
		monkey.Operation = func(left, right float64) float64 {
			return left / right
		}
	}
}

func CreateCalculatingMonkey(input string) *CalculationMonkey {
	monkey := new(CalculationMonkey)

	parts := strings.Split(input, ":")
	value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err == nil {
		monkey.Value = float64(value)
	} else {
		MonkeyOperationFromInput(monkey, parts[1])
	}

	monkey.Name = parts[0]

	return monkey
}

func AddWaitingMonkeys(monkey *CalculationMonkey, waitingOn map[string][]*CalculationMonkey) {
	for _, name := range monkey.WaitingOn {
		waitingMonkeys, ok := waitingOn[name]
		if !ok {
			waitingMonkeys = []*CalculationMonkey{}
		}
		waitingMonkeys = append(waitingMonkeys, monkey)

		waitingOn[name] = waitingMonkeys
	}
}

func RemoveWaitingFromMonkey(monkey *CalculationMonkey, name string) {
	left := []string{}

	for _, waitingName := range monkey.WaitingOn {
		if waitingName != name {
			left = append(left, waitingName)
		}
	}

	monkey.WaitingOn = left
}

func ResolveWaitingMonkeys(initialMonkey *CalculationMonkey, monkeysByName map[string]*CalculationMonkey) {
	waitingOn := make(map[string][]*CalculationMonkey)

	AddWaitingMonkeys(initialMonkey, waitingOn)
	for len(initialMonkey.WaitingOn) > 0 {
		readyToCalculate := []*CalculationMonkey{}

		for key, waitingMonkeys := range waitingOn {
			monkey := monkeysByName[key]
			if len(monkey.WaitingOn) > 0 {
				AddWaitingMonkeys(monkey, waitingOn)
			} else {
				for _, waiting := range waitingMonkeys {
					RemoveWaitingFromMonkey(waiting, key)
					if len(waiting.WaitingOn) == 0 {
						readyToCalculate = append(readyToCalculate, waiting)
					}
				}
			}
		}

		for _, monkey := range readyToCalculate {
			left := monkeysByName[monkey.Dependants[0]]
			right := monkeysByName[monkey.Dependants[1]]

			monkey.Value = monkey.Operation(left.Value, right.Value)
		}
	}
}

func DayTwentyoneFirst(monkeysByName map[string]*CalculationMonkey) int {
	initial := monkeysByName["root"]

	ResolveWaitingMonkeys(initial, monkeysByName)

	return int(initial.Value)
}

func DayTwentyoneSecond(monkeysByName map[string]*CalculationMonkey) int {
	initial := monkeysByName["root"]
	initial.Value = 2000
	initial.Operation = func(f1, f2 float64) float64 {
		return f1 - f2
	}

	human := monkeysByName["humn"]

	human.Value = 3296135418819
	var current float64 = 3296135418819
	for i := 1; i < 20; i++ {
		ResolveWaitingMonkeys(initial, monkeysByName)

		if initial.Value == 0 {
			return int(current)
		}

		if initial.Value > 0 {
			current = current + float64(i)
		}
		for _, cm := range monkeysByName {
			cm.Reset()
		}
		human.Value = current
	}

	// if waiting on humn
	//  add to a list with monkeys to evaluate at the end
	//  replace operation with inverse
	// if humn: skip
	return 0
}

func DayTwentyone() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d21.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	monkeysByName := make(map[string]*CalculationMonkey)
	monkeysByNameB := make(map[string]*CalculationMonkey)
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		monkey := CreateCalculatingMonkey(input)
		monkeyB := CreateCalculatingMonkey(input)

		monkeysByName[monkey.Name] = monkey
		monkeysByNameB[monkeyB.Name] = monkeyB
	}

	first := DayTwentyoneFirst(monkeysByName)

	second := DayTwentyoneSecond(monkeysByNameB)

	openFile.File.Close()
	return first, second
}
