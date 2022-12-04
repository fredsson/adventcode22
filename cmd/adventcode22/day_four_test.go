package main

import (
	"fmt"
	"testing"
)

func Test_CreateAssignmentGroup_shouldGiveCorrectGroups(t *testing.T) {
	assignmentGroup := CreateAssignmentGroup("1-4,5-23")

	if assignmentGroup.First.Start != 1 && assignmentGroup.First.End != 4 {
		t.Error("Incorrect first group")
	}

	if assignmentGroup.Second.Start != 5 && assignmentGroup.Second.End != 23 {
		t.Error("Incorrect second group")
	}
}

func Test_AssignmentRangeContainsOtherRange_shouldGiveCorrectAnswer(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"1-2,3-4", false},
		{"1-4,2-3", true},
		{"2-3,1-4", true},
		{"2-3,1-3", true},
		{"2-3,2-5", true},
		{"1-2,1-2", true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("for %s", tc.input), func(t *testing.T) {
			group := CreateAssignmentGroup(tc.input)

			result := AssignmentRangeContainsOtherRange(group)

			if result != tc.want {
				t.Errorf("result is incorrect, is %t wanted %t for %s", result, tc.want, tc.input)
			}

		})
	}
}

func Test_AssignmentRangeOverlap_shouldGiveCorrectAnswer(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{"1-2,3-4", false},
		{"5-7,7-9", true},
		{"16-21,34-70", false},
		{"28-81,7-27", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("for %s", tc.input), func(t *testing.T) {
			group := CreateAssignmentGroup(tc.input)

			result := AssignmentRangeOverlap(group)

			if result != tc.want {
				t.Errorf("result is incorrect, is %t wanted %t for %s", result, tc.want, tc.input)
			}

		})
	}
}
