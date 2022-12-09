package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AssignmentRange struct {
	Start int
	End   int
}

type AssignmentGroup struct {
	First  AssignmentRange
	Second AssignmentRange
}

func parseAssignmentRangeFromString(group string) AssignmentRange {
	groupRange := strings.Split(group, "-")

	start, err := strconv.Atoi(groupRange[0])
	if err != nil {
		fmt.Println("Failed to parse assignment to number", groupRange[0])
		return AssignmentRange{
			Start: 0,
			End:   0,
		}
	}
	end, err := strconv.Atoi(groupRange[1])
	if err != nil {
		fmt.Println("Failed to parse assignment to number", groupRange[1])
		return AssignmentRange{
			Start: 0,
			End:   0,
		}
	}

	return AssignmentRange{
		Start: start,
		End:   end,
	}
}

func CreateAssignmentGroup(input string) *AssignmentGroup {
	groups := strings.Split(input, ",")
	assignmentGroup := new(AssignmentGroup)
	assignmentGroup.First = parseAssignmentRangeFromString(groups[0])
	assignmentGroup.Second = parseAssignmentRangeFromString(groups[1])
	return assignmentGroup
}

func AssignmentRangeContainsOtherRange(group *AssignmentGroup) bool {
	if group.First.Start <= group.Second.Start && group.First.End >= group.Second.End {
		return true
	}
	if group.Second.Start <= group.First.Start && group.Second.End >= group.First.End {
		return true
	}
	return false
}

func AssignmentRangeOverlap(group *AssignmentGroup) bool {
	if group.First.Start > group.Second.End {
		return false
	}
	if group.First.End >= group.Second.Start {
		return true
	}
	return false
}

func DayFour() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d4.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	groupsForReconsideration := 0
	groupsWithOverlap := 0
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()
		group := CreateAssignmentGroup(input)

		if AssignmentRangeContainsOtherRange(group) {
			groupsForReconsideration++
		}

		if AssignmentRangeOverlap(group) {
			groupsWithOverlap++
		}
	}

	openFile.File.Close()
	return groupsForReconsideration, groupsWithOverlap
}
