package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RockLine struct {
	Points []Location
}

func CreateRockLine(input string) *RockLine {
	line := new(RockLine)

	for _, v := range strings.Split(input, "->") {
		trimmed := strings.TrimSpace(v)

		coordinates := strings.Split(trimmed, ",")

		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])

		line.Points = append(line.Points, Location{x, y})
	}

	return line
}

func SimulateSandNeededToReachBottom(world []string, max Location) bool {
	currentPosition := Location{500, 0}

	hasMoved := true
	for hasMoved {
		if currentPosition.y >= max.y-2 {
			return true
		}

		down := ((currentPosition.y + 1) * max.x) + currentPosition.x
		downLeft := ((currentPosition.y + 1) * max.x) + currentPosition.x - 1
		downRight := ((currentPosition.y + 1) * max.x) + currentPosition.x + 1

		if world[down] == "." {
			currentPosition = Location{currentPosition.x, currentPosition.y + 1}
		} else if world[downLeft] == "." {
			currentPosition = Location{currentPosition.x - 1, currentPosition.y + 1}
		} else if world[downRight] == "." {
			currentPosition = Location{currentPosition.x + 1, currentPosition.y + 1}
		} else {
			hasMoved = false
		}
	}

	index := (currentPosition.y * max.x) + currentPosition.x
	world[index] = "@"

	return false
}

func SimulateSandNeededToReachStart(world []string, max Location) bool {
	currentPosition := Location{500, 0}

	hasMoved := true
	for hasMoved {
		down := ((currentPosition.y + 1) * max.x) + currentPosition.x
		downLeft := ((currentPosition.y + 1) * max.x) + currentPosition.x - 1
		downRight := ((currentPosition.y + 1) * max.x) + currentPosition.x + 1

		if world[down] == "." {
			currentPosition = Location{currentPosition.x, currentPosition.y + 1}
		} else if world[downLeft] == "." {
			currentPosition = Location{currentPosition.x - 1, currentPosition.y + 1}
		} else if world[downRight] == "." {
			currentPosition = Location{currentPosition.x + 1, currentPosition.y + 1}
		} else {
			hasMoved = false
		}
	}

	index := (currentPosition.y * max.x) + currentPosition.x
	world[index] = "@"

	if currentPosition.x == 500 && currentPosition.y == 0 {
		return true
	}

	return false

}

func FindMax(lines []*RockLine) Location {
	max := Location{math.MinInt32, 0}

	for _, line := range lines {
		for _, point := range line.Points {
			if point.x > max.x {
				max.x = point.x + 1
			}
			if point.y > max.y {
				max.y = point.y + 1
			}
		}
	}

	return max
}

func DrawLine(line *RockLine, world []string, width int) {
	for index := range line.Points {
		if index > 0 {
			previous := line.Points[index-1]
			current := line.Points[index]

			startx := int(math.Min(float64(previous.x), float64(current.x)))
			endx := int(math.Max(float64(previous.x), float64(current.x)))
			startY := int(math.Min(float64(previous.y), float64(current.y)))
			endY := int(math.Max(float64(previous.y), float64(current.y)))

			for y := startY; y <= endY; y++ {
				for x := startx; x <= endx; x++ {
					worldIndex := (y * width) + x
					world[worldIndex] = "#"
				}
			}
		}
	}
}

func DrawWorld(world []string, max Location) {
	for y := 0; y < max.y; y++ {
		for x := 0; x < max.x; x++ {
			i := (y * max.x) + x
			fmt.Print(world[i])
		}
		fmt.Println()
	}
}

func DayFourteen() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d14.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	rockLines := []*RockLine{}
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		line := CreateRockLine(input)
		rockLines = append(rockLines, line)
	}

	max := FindMax(rockLines)
	max.x = 1000
	max.y = max.y + 2

	world := []string{}

	for y := 0; y < max.y; y++ {
		for x := 0; x < max.x; x++ {
			world = append(world, ".")
		}
	}

	for _, line := range rockLines {
		DrawLine(line, world, max.x)
	}

	bottomLine := CreateRockLine(fmt.Sprintf("0,%d -> 999,%d", max.y-1, max.y-1))
	DrawLine(bottomLine, world, max.x)

	SandHasReachedAbyss := false
	unitsOfSand := 0
	for !SandHasReachedAbyss {
		SandHasReachedAbyss = SimulateSandNeededToReachBottom(world, max)
		if !SandHasReachedAbyss {
			unitsOfSand++
		}
	}

	SandHasReachedTop := false
	unitsOfSandB := unitsOfSand + 1
	for !SandHasReachedTop {
		SandHasReachedTop = SimulateSandNeededToReachStart(world, max)
		if !SandHasReachedTop {
			unitsOfSandB++
		}
	}

	openFile.File.Close()
	return unitsOfSand, unitsOfSandB
}
