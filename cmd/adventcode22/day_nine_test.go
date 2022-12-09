package main

import (
	"fmt"
	"testing"
)

func Test_TailNeedsMoving(t *testing.T) {
	testCases := []struct {
		knotCoordinates Coordinates
		tailCoordinates Coordinates
		want            bool
	}{
		{Coordinates{0, 0}, Coordinates{0, 0}, false},
		{Coordinates{2, 0}, Coordinates{0, 0}, true},
		{Coordinates{0, 2}, Coordinates{0, 0}, true},
		{Coordinates{2, 2}, Coordinates{0, 0}, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("for %+v %+v", tc.knotCoordinates, tc.tailCoordinates), func(t *testing.T) {
			knot := CreateKnot()

			knot.Position = tc.knotCoordinates
			knot.TailPosition = tc.tailCoordinates

			needsMoving := TailNeedsMoving(knot)

			if needsMoving != tc.want {
				t.Error("Needs moving incorrect expected", tc.want, "but got", needsMoving)
			}
		})
	}
}
