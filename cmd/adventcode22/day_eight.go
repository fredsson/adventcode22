package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ShadowedByLargerTree(index int, patch []string) bool {
	height, _ := strconv.Atoi(patch[index])

	x := index % 99
	y := (index / 99)

	visibleInNegativeX := true
	for i := x - 1; i >= 0; i-- {
		newIndex := (y * 99) + i
		compare, _ := strconv.Atoi(patch[newIndex])
		if compare >= height {
			visibleInNegativeX = false
			break
		}
	}

	visibleInPositiveX := true
	for i := x + 1; i < 99; i++ {
		newIndex := (y * 99) + i
		compare, _ := strconv.Atoi(patch[newIndex])
		if compare >= height {
			visibleInPositiveX = false
			break
		}
	}

	visibleInNegativeY := true
	for i := y - 1; i >= 0; i-- {
		newIndex := (i * 99) + x
		compare, _ := strconv.Atoi(patch[newIndex])
		if compare >= height {
			visibleInNegativeY = false
			break
		}
	}

	visibleInPositiveY := true
	for i := y + 1; i < 99; i++ {
		newIndex := (i * 99) + x
		compare, _ := strconv.Atoi(patch[newIndex])
		if compare >= height {
			visibleInPositiveY = false
			break
		}
	}

	visibleInX := visibleInNegativeX || visibleInPositiveX
	visibleInY := visibleInNegativeY || visibleInPositiveY
	return !(visibleInX || visibleInY)
}

func CalculateScenicScore(index int, patch []string) int {
	height, _ := strconv.Atoi(patch[index])

	x := index % 99
	y := (index / 99)

	visibleRangeInNegativeX := 0
	for i := x - 1; i >= 0; i-- {
		newIndex := (y * 99) + i
		compare, _ := strconv.Atoi(patch[newIndex])
		visibleRangeInNegativeX++
		if compare >= height {
			break
		}
	}

	visibleRangeInPositiveX := 0
	for i := x + 1; i < 99; i++ {
		newIndex := (y * 99) + i
		compare, _ := strconv.Atoi(patch[newIndex])
		visibleRangeInPositiveX++
		if compare >= height {
			break
		}
	}

	visibleRangeInNegativeY := 0
	for i := y - 1; i >= 0; i-- {
		newIndex := (i * 99) + x
		compare, _ := strconv.Atoi(patch[newIndex])
		visibleRangeInNegativeY++
		if compare >= height {
			break
		}
	}

	visibleRangeInPositiveY := 0
	for i := y + 1; i < 99; i++ {
		newIndex := (i * 99) + x
		compare, _ := strconv.Atoi(patch[newIndex])
		visibleRangeInPositiveY++
		if compare >= height {
			break
		}
	}

	return (visibleRangeInNegativeX * visibleRangeInPositiveX * visibleRangeInNegativeY * visibleRangeInPositiveY)
}

func DayEight() {
	openFile := readFileByLines("inputs/d8.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return
	}

	treePatch := []string{}
	patchWidth := 0
	patchHeight := 0
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()
		patchWidth = len(input)
		patchHeight++
		treePatch = append(treePatch, strings.Split(input, "")...)
	}

	visibleTrees := 0
	highestScore := 0
	for index := 0; index < patchWidth*patchHeight; index++ {
		scenicScore := CalculateScenicScore(index, treePatch)
		if highestScore < scenicScore {
			highestScore = scenicScore
		}
		if index < patchWidth {
			visibleTrees++
			continue
		}
		if index%patchWidth == 0 || index%patchWidth == patchWidth-1 {
			visibleTrees++
			continue
		}
		if index >= (patchWidth*patchHeight)-patchWidth {
			visibleTrees++
			continue
		}

		if !ShadowedByLargerTree(index, treePatch) {
			visibleTrees++
		}
	}

	fmt.Println(visibleTrees)
	fmt.Println(highestScore)

	openFile.File.Close()
}
