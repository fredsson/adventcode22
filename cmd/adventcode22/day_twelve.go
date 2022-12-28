package main

import (
	"fmt"
	"math"
)

type Tile struct {
	index, x, y, height int
}

func (t *Tile) Neighbors(width, height int) []int {
	possible := []struct {
		x int
		y int
	}{
		{1, 0},
		{-1, 0},
		{0, -1},
		{0, 1},
	}

	result := []int{}
	for _, d := range possible {
		ny := t.y + d.y
		nx := t.x + d.x
		if nx < 0 || nx >= width {
			continue
		}
		if ny < 0 || ny >= height {
			continue
		}

		nindex := (ny * width) + nx
		result = append(result, nindex)
	}

	return result
}

func (t *Tile) CanClimbWithoutGear(other *Tile) bool {
	return (other.height - t.height) <= 1
}

func CreateTile(input rune, width, x, y int) *Tile {
	t := new(Tile)

	// S and E can be lowest possible
	if input == 83 {
		t.height = 0
	} else if input == 69 {
		t.height = (122 - 97)
	} else {
		t.height = (int(input) - 97)
	}
	t.index = (y * width) + x
	t.x = x
	t.y = y

	return t
}

func BuildFrontier(from *Tile, tiles []*Tile, width, height int) map[int]int {
	frontier := []*Tile{}
	frontier = append(frontier, from)

	paths := make(map[int]int)
	paths[from.index] = -1

	var current *Tile
	for len(frontier) > 0 {
		current, frontier = frontier[0], frontier[1:]

		indices := current.Neighbors(width, height)
		for _, n := range indices {
			neighbor := tiles[n]

			if !current.CanClimbWithoutGear(neighbor) {

				continue
			}

			_, ok := paths[n]
			if ok {
				continue
			}

			frontier = append(frontier, neighbor)
			paths[neighbor.index] = current.index
		}
	}

	return paths
}

func FindShortestPath(from *Tile, to *Tile, tiles []*Tile, width, height int) []*Tile {
	frontier := BuildFrontier(from, tiles, width, height)

	path := []*Tile{}
	current := to
	didNotFoundPath := false
	for current != nil && current.index != from.index {
		path = append(path, current)
		nextIndex, ok := frontier[current.index]
		if !ok {
			didNotFoundPath = true
		}
		if !ok || nextIndex == -1 {
			break
		}

		current = tiles[nextIndex]
	}

	if didNotFoundPath {
		return []*Tile{}
	}

	return path
}

func DayTwelve() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d12.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	tiles := []*Tile{}
	var width int
	var height int = 0
	var start *Tile
	var end *Tile
	starts := []*Tile{}
	for openFile.Scanner.Scan() {
		line := openFile.Scanner.Text()
		width = len(line)
		for x, tile := range line {
			t := CreateTile(tile, width, x, height)
			if string(tile) == "S" {
				start = t
			}
			if string(tile) == "E" {
				end = t
			}

			if string(tile) == "a" {
				starts = append(starts, t)
			}

			tiles = append(tiles, t)
		}
		height++
	}

	path := FindShortestPath(start, end, tiles, width, height)

	shortestPathAmongStarts := math.MaxInt32

	for _, s := range starts {
		possiblePath := FindShortestPath(s, end, tiles, width, height)
		if len(possiblePath) > 0 {
			shortestPathAmongStarts = int(math.Min(float64(shortestPathAmongStarts), float64(len(possiblePath))))
		}
	}

	openFile.File.Close()

	return len(path), shortestPathAmongStarts
}
