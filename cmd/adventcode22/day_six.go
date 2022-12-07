package main

import "fmt"

func areValuesUnique(values []string) bool {
	uniques := make(map[string]bool)
	for _, value := range values {
		if uniques[value] {
			return false
		}
		uniques[value] = true
	}
	return true
}

func DaySix() {
	openFile := readFileByLines("inputs/d6.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return
	}

	openFile.Scanner.Scan()
	input := openFile.Scanner.Text()

	buffer := []string{}

	for index, code := range input {
		char := string(code)
		if len(buffer) < 4 {
			buffer = append(buffer, char)
		} else {
			buffer = append(buffer[1:], char)
		}
		if len(buffer) >= 4 {
			if areValuesUnique(buffer) {
				fmt.Println(index + 1)
				break
			}
		}
	}

	bufferB := []string{}
	for index, code := range input {
		char := string(code)
		if len(bufferB) < 14 {
			bufferB = append(bufferB, char)
		} else {
			bufferB = append(bufferB[1:], char)
		}
		if len(bufferB) >= 14 {
			if areValuesUnique(bufferB) {
				fmt.Println(index + 1)
				break
			}
		}
	}

	openFile.File.Close()
}
