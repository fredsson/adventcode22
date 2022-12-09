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
	Body               []Coordinates
	TailPosition       Coordinates
	TailVisitedByIndex map[int]bool
}

func CreateKnot() *Knot {
	knot := new(Knot)

	knot.Position = Coordinates{}
	knot.PreviousPosition = Coordinates{}
	knot.Body = []Coordinates{}
	knot.TailPosition = Coordinates{}
	knot.TailVisitedByIndex = make(map[int]bool)
	knot.TailVisitedByIndex[0] = true

	return knot
}

func CreateMultiPartKnot(size int) *Knot {
	knot := new(Knot)

	knot.Position = Coordinates{}
	knot.PreviousPosition = Coordinates{}
	knot.Body = []Coordinates{}
	for i := 0; i < size; i++ {
		knot.Body = append(knot.Body, Coordinates{})
	}

	knot.TailPosition = Coordinates{}
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

func CoordinatesAdjecent(a Coordinates, b Coordinates) bool {
	dx := math.Abs(float64(a.x - b.x))
	dy := math.Abs(float64(a.y - b.y))

	return dx <= 1 && dy <= 1
}

func MakeCoordinatesAdjecent(current Coordinates, target Coordinates) Coordinates {
	x := 0
	if current.x < target.x {
		x = 1
	} else if current.x > target.x {
		x = -1
	}
	y := 0
	if current.y < target.y {
		y = 1
	} else if current.y > target.y {
		y = -1
	}
	return Coordinates{current.x + x, current.y + y}
}

func (knot *Knot) MoveHead(direction Coordinates) {
	knot.PreviousPosition = knot.Position
	knot.Position = Coordinates{
		knot.Position.x + direction.x,
		knot.Position.y + direction.y,
	}
}

func (knot *Knot) MoveBody() {
	previous := knot.Position
	for index, body := range knot.Body {
		if !CoordinatesAdjecent(previous, body) {
			knot.Body[index] = MakeCoordinatesAdjecent(body, previous)
		}

		previous = knot.Body[index]
	}
}

func (knot *Knot) MoveTail() {
	previous := knot.Position
	if len(knot.Body) > 0 {
		previous = knot.Body[len(knot.Body)-1]
	}
	if !CoordinatesAdjecent(previous, knot.TailPosition) {
		knot.TailPosition = MakeCoordinatesAdjecent(knot.TailPosition, previous)
		index := (knot.TailPosition.y * 1000) + knot.TailPosition.x
		knot.TailVisitedByIndex[index] = true
	}
}

func (knot *Knot) Move(direction string, steps int) {
	dir := CoordinatesFromDirection(direction)
	for i := 0; i < steps; i++ {
		knot.MoveHead(dir)
		knot.MoveBody()
		knot.MoveTail()
	}
}

func DayNine() {
	openFile := readFileByLines("inputs/d9.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return
	}

	knot := CreateKnot()
	multiKnot := CreateMultiPartKnot(8)
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		commands := strings.Split(input, " ")
		steps, _ := strconv.Atoi(commands[1])
		knot.Move(commands[0], steps)
		multiKnot.Move(commands[0], steps)
	}

	fmt.Println(len(knot.TailVisitedByIndex))
	fmt.Println(len(multiKnot.TailVisitedByIndex))

	openFile.File.Close()
}
