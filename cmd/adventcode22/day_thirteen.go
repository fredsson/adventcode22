package main

import "fmt"

func DayThirteen() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d13.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	for openFile.Scanner.Scan() {
		// tokenize line?
		// build AST
		// compare left tree with right tree
		// if same type
		//  compare type
		//  if left < right
		//    return true
		// if different types
		//  convert to same type
		//  compare type
		//  if left < right
		//    return true
		//
		//  return false
	}

	openFile.File.Close()
	return 0, 0
}
