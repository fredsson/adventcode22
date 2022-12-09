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
	Head               Coordinates
	Body               []Coordinates
	Tail               Coordinates
	TailVisitedByIndex map[int]bool
}

func CreateKnot() *Knot {
	knot := new(Knot)

	knot.Head = Coordinates{}
	knot.Body = []Coordinates{}
	knot.Tail = Coordinates{}
	knot.TailVisitedByIndex = make(map[int]bool)
	knot.TailVisitedByIndex[0] = true

	return knot
}

func CreateMultiPartKnot(size int) *Knot {
	knot := new(Knot)

	knot.Head = Coordinates{}
	knot.Body = []Coordinates{}
	for i := 0; i < size; i++ {
		knot.Body = append(knot.Body, Coordinates{})
	}

	knot.Tail = Coordinates{}
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
	knot.Head = Coordinates{
		knot.Head.x + direction.x,
		knot.Head.y + direction.y,
	}
}

func (knot *Knot) MoveBody() {
	previous := knot.Head
	for index, body := range knot.Body {
		if !CoordinatesAdjecent(previous, body) {
			knot.Body[index] = MakeCoordinatesAdjecent(body, previous)
		}

		previous = knot.Body[index]
	}
}

func (knot *Knot) MoveTail() {
	previous := knot.Head
	if len(knot.Body) > 0 {
		previous = knot.Body[len(knot.Body)-1]
	}
	if !CoordinatesAdjecent(previous, knot.Tail) {
		knot.Tail = MakeCoordinatesAdjecent(knot.Tail, previous)
		index := (knot.Tail.y * 1000) + knot.Tail.x
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
