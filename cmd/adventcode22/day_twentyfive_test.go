package main

import (
	"fmt"
	"testing"
)

func Test_SnafuToDecimalConverter(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"1", 1},
		{"2", 2},
		{"1=", 3},
		{"1-", 4},
		{"10", 5},
		{"11", 6},
		{"2=", 8},
		{"20", 10},
		{"1=0", 15},
		{"21", 11},
		{"34", 19},
		{"1-0", 20},
		{"100", 25},
		{"111", 31},
		{"1=11-2", 2022},
		{"1-0---0", 12345},
		{"1121-1110-1=0", 314159265},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("for %s", tc.input), func(t *testing.T) {
			decimal := SnafuToDecimal(tc.input)

			if decimal != tc.want {
				t.Error("Did not convert snafu correctly, expected", tc.want, "received", decimal)
			}
		})
	}
}

func Test_DecimalToSnafuConverter(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{1, "1"},
		{2, "2"},
		{3, "1="},
		{4, "1-"},
		{5, "10"},
		{6, "11"},
		{8, "2="},
		{10, "20"},
		{11, "21"},
		{15, "1=0"},
		{20, "1-0"},
		{25, "100"},
		{107, "1-12"},
		{107, "1-12"},
		{37, "122"},
		{353, "1=-1="},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("for %d", tc.input), func(t *testing.T) {
			snafu := DecimalToSnafu(tc.input)

			if snafu != tc.want {
				t.Error("Did not convert snafu correctly, expected", tc.want, "received", snafu)
			}
		})
	}

}
