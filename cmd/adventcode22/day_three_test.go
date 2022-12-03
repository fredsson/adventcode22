package main

import (
	"fmt"
	"testing"
)

func Test_createRucksack_shouldGiveCorrectCompartments(t *testing.T) {
	rucksack := CreateRucksack("abAB")

	if rucksack.LeftCompartment != "ab" {
		t.Error("Incorrect left compartment")
	}
	if rucksack.RightCompartment != "AB" {
		t.Error("Incorrect right compartment")
	}
}

func Test_priorityOfMissplacedItems(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"abAB", 0},
		{"abaB", 1},
		{"abAb", 2},
		{"acAc", 3},
		{"jCAj", 10},
		{"aBaB", 29},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("for %s", tc.input), func(t *testing.T) {

			rucksack := CreateRucksack(tc.input)

			priority := rucksack.PriorityOfMissplacedItems()

			if priority != tc.want {
				t.Errorf("Priority is incorrect, is %d wanted %d for %s", priority, tc.want, tc.input)
			}
		})
	}

}

func Test_priorityOfMissplacedItems_withMissplaced_shouldReturnPriority(t *testing.T) {
	rucksack := CreateRucksack("abaB")

	priority := rucksack.PriorityOfMissplacedItems()

	if priority != 1 {
		t.Error("Priority is incorrect")
	}
}
