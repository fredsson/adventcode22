package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coordinates struct {
	x int
	y int
}

type Knot struct {
	Position           Coordinates
	PreviousPosition   Coordinates
	TailPosition       Coordinates
	TailVisitedByIndex map[int]bool
}

func CreateKnot() *Knot {
	knot := new(Knot)

	knot.Position = Coordinates{0, 0}
	knot.PreviousPosition = Coordinates{0, 0}
	knot.TailPosition = Coordinates{0, 0}
	knot.TailVisitedByIndex = make(map[int]bool)
	knot.TailVisitedByIndex[0] = true

	return knot
}

func CoordinatesFromDirection(direction string) Coordinates {
	switch direction {
	case "L":
		return Coordinates{-1, 0}
	case "R":
		return Coordinates{1, 0}
	case "U":
		return Coordinates{0, -1}
	case "D":
		return Coordinates{0, 1}
	}
	return Coordinates{0, 0}
}

func TailNeedsMoving(knot *Knot) bool {
	if math.Abs(float64(knot.Position.x-knot.TailPosition.x)) > 1 {
		return true
	}
	if math.Abs(float64(knot.Position.y-knot.TailPosition.y)) > 1 {
		return true
	}
	return false
}

func (knot *Knot) Move(direction string, steps int) {
	dir := CoordinatesFromDirection(direction)
	for i := 0; i < steps; i++ {
		knot.PreviousPosition = knot.Position
		knot.Position.x += dir.x
		knot.Position.y += dir.y
		if TailNeedsMoving(knot) {
			knot.TailPosition = knot.PreviousPosition
			index := (knot.TailPosition.y * 1000) + knot.TailPosition.x
			knot.TailVisitedByIndex[index] = true
		}
	}
}

func DayNine() {
	openFile := readFileByLines("inputs/d9.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return
	}

	knot := CreateKnot()
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		commands := strings.Split(input, " ")
		steps, _ := strconv.Atoi(commands[1])
		knot.Move(commands[0], steps)
	}

	fmt.Println(len(knot.TailVisitedByIndex))

	openFile.File.Close()
}
