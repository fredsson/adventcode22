package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func DecimalToSnafu(decimal int) string {
	// since snafu is base 5 but in range [-2, 3]
	snafuOffset := 2

	snafu := ""

	v := decimal
	for v != 0 {
		reminder := (v + snafuOffset) % 5
		tokenValue := reminder - snafuOffset

		var snafuToken string
		if tokenValue == -2 {
			snafuToken = "="
		} else if tokenValue == -1 {
			snafuToken = "-"
		} else {
			snafuToken = fmt.Sprintf("%d", tokenValue)
		}
		snafu += snafuToken

		v = (v - tokenValue) / 5
	}

	return reverse(snafu)
}

func SnafuToDecimal(snafu string) int {
	parts := strings.Split(snafu, "")
	result := 0

	for i := len(parts) - 1; i >= 0; i-- {
		fivePow := int(math.Pow(float64(5), float64(len(parts)-1-i)))
		token := parts[i]
		value := 0
		amount := 0
		if token == "=" {
			value = -2 * fivePow
			amount = 1
		} else if token == "-" {
			value = -1 * fivePow
			amount = 1
		} else {
			amount, _ = strconv.Atoi(parts[i])
			value = fivePow
		}
		if value == 0 {
			result += amount
		} else {
			result += amount * value
		}
	}

	return result
}

func DayTwentyfive() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d25.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	sum := 0
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		decimal := SnafuToDecimal(input)

		sum += decimal
	}

	openFile.File.Close()
	return DecimalToSnafu(sum), 0
}
