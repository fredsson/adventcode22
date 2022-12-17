package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

type Location struct {
	x int
	y int
}

type Sensor struct {
	location      Location
	closestBeacon Location
}

type OccupiedSlice struct {
	start int
	end   int
}

type PossibleSlice struct {
	y      int
	slices []OccupiedSlice
}

func ManhattanDistance(a Location, b Location) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func CalculateSlice(beacon Location, sensor Location, y int) (OccupiedSlice, error) {
	opposite := int(math.Abs(float64(y - sensor.y)))
	hypotenuse := ManhattanDistance(beacon, sensor)

	if opposite > hypotenuse {
		return OccupiedSlice{}, errors.New("not In Range")
	}

	adjacent := int(math.Abs(float64(hypotenuse - opposite)))

	left := sensor.x - adjacent
	right := sensor.x + adjacent

	return OccupiedSlice{left, right}, nil
}

func CollapseSlices(input []OccupiedSlice) []OccupiedSlice {
	result := []OccupiedSlice{}

	for _, slice := range input {

		updated := false
		for i, existingSlice := range result {
			if slice.end > existingSlice.end && slice.start <= existingSlice.end {
				result[i].end = slice.end
				updated = true
			} else if slice.end <= existingSlice.end && slice.start >= existingSlice.start {
				updated = true
			} else if slice.start < existingSlice.start && slice.end >= existingSlice.start {
				result[i].start = slice.start
				updated = true
			}
		}
		if !updated {
			result = append(result, slice)
		}
	}

	return result
}

func CalculateSliceLength(slice OccupiedSlice) int {
	return slice.end - slice.start
}

func FirstPart(sensors []Sensor) int {
	occupiedSlices := []OccupiedSlice{}
	for _, sensor := range sensors {
		slice, err := CalculateSlice(sensor.closestBeacon, sensor.location, 2_000_000)
		if err == nil {
			occupiedSlices = append(occupiedSlices, slice)
		}
	}
	sort.SliceStable(occupiedSlices, func(i, j int) bool {
		return occupiedSlices[i].start < occupiedSlices[j].start
	})

	collapsedSlices := CollapseSlices(occupiedSlices)

	totalOccupied := 0
	for _, collapsed := range collapsedSlices {
		totalOccupied += CalculateSliceLength(collapsed)
	}

	return totalOccupied
}

func SecondPart(sensors []Sensor) int {

	possibleSlices := []PossibleSlice{}
	for y := 0; y < 4_000_000; y++ {
		occupiedSlices := []OccupiedSlice{}
		for _, sensor := range sensors {
			slice, err := CalculateSlice(sensor.closestBeacon, sensor.location, y)
			if err == nil {
				occupiedSlices = append(occupiedSlices, slice)
			}
		}
		sort.SliceStable(occupiedSlices, func(i, j int) bool {
			return occupiedSlices[i].start < occupiedSlices[j].start
		})

		collapsedSlices := CollapseSlices(occupiedSlices)
		if len(collapsedSlices) > 1 {
			possibleSlices = append(possibleSlices, PossibleSlice{y, collapsedSlices})
		}
	}

	x := possibleSlices[0].slices[0].end + 1
	y := possibleSlices[0].y

	return x*4_000_000 + y
}

func DayFifteen() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d15.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	sensors := []Sensor{}
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		location := Location{0, 0}
		beacon := Location{0, 0}
		fmt.Sscanf(input, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &location.x, &location.y, &beacon.x, &beacon.y)

		sensors = append(sensors, Sensor{location, beacon})
	}

	totalOccupied := FirstPart(sensors)

	tuningFrequency := SecondPart(sensors)

	openFile.File.Close()
	return totalOccupied, tuningFrequency
}
