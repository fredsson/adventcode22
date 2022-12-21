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

func NamesFromOperation(operation string, operator string) []string {
	result := []string{}
	names := strings.Split(operation, operator)
	for _, name := range names {
		result = append(result, strings.TrimSpace(name))
	}

	return result
}

func MonkeyOperationFromInput(input string) ([]string, OperationFunc) {
	var result []string
	var operation OperationFunc

	if strings.Contains(input, "*") {
		result = NamesFromOperation(input, "*")
		operation = func(left float64, right float64) float64 {
			return left * right
		}
	} else if strings.Contains(input, "-") {
		result = NamesFromOperation(input, "-")
		operation = func(left float64, right float64) float64 {
			return left - right
		}
	} else if strings.Contains(input, "+") {
		result = NamesFromOperation(input, "+")
		operation = func(left float64, right float64) float64 {
			return left + right
		}
	} else if strings.Contains(input, "/") {
		result = NamesFromOperation(input, "/")
		operation = func(left float64, right float64) float64 {
			return left / right
		}
	}

	return result, operation
}

func CreateCalculatingMonkey(input string) *CalculationMonkey {
	monkey := new(CalculationMonkey)

	parts := strings.Split(input, ":")
	value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err == nil {
		monkey.Value = float64(value)
	} else {
		names, operation := MonkeyOperationFromInput(parts[1])
		monkey.WaitingOn = names
		monkey.Dependants = names
		monkey.Operation = operation
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

func DayTwentyone() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d21.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	monkeysByName := make(map[string]*CalculationMonkey)
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		monkey := CreateCalculatingMonkey(input)

		monkeysByName[monkey.Name] = monkey
	}

	initial := monkeysByName["root"]

	waitingOn := make(map[string][]*CalculationMonkey)

	AddWaitingMonkeys(initial, waitingOn)

	for len(initial.WaitingOn) > 0 {
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

	openFile.File.Close()
	return int(initial.Value), 0
}
