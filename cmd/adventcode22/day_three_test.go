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

func Test_PriorityOfMissplacedItems(t *testing.T) {
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

func Test_FindPriorityForMatchingItem(t *testing.T) {
	testCases := []struct {
		rucksack1 string
		rucksack2 string
		rucksack3 string
		want      int
	}{
		{"HjAB", "abCD", "abDE", 0},
		{"abAB", "abCD", "acDE", 1},
		{"HDAB", "abCD", "abDE", 30},
		{"HdZB", "abZD", "ZbDE", 52},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("for %s, %s, %s", tc.rucksack1, tc.rucksack2, tc.rucksack3), func(t *testing.T) {

			rucksack1 := CreateRucksack(tc.rucksack1)
			rucksack2 := CreateRucksack(tc.rucksack2)
			rucksack3 := CreateRucksack(tc.rucksack3)

			priority := rucksack1.FindPriorityForMatchingItem(rucksack2, rucksack3)

			if priority != tc.want {
				t.Errorf("Priority is incorrect, is %d wanted %d for %s", priority, tc.want, tc.rucksack1)
			}
		})
	}
}
