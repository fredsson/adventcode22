package main

import "testing"

func Test_MoveSupply_ShouldMoveSingleSupply(t *testing.T) {
	supplyArea := CreateSupplyArea([]string{
		"[H] [B] [R] [S] [R] [T] [S] [R] [L]",
		" 1   2   3   4   5   6   7   8   9 ",
	})

	supplyArea.MoveSupply(1, 1, 3)

	if len(supplyArea.areas[0]) != 0 {
		t.Error("Did not remove supply from correct area")
	}

	if len(supplyArea.areas[2]) != 2 {
		t.Error("Did not add supply to correct area")
	}

	if supplyArea.areas[2][1] != "H" {
		t.Error("Did not add supply to correct area")
	}
}

func Test_MoveSupply_ShouldMoveMultipleSupply(t *testing.T) {
	supplyArea := CreateSupplyArea([]string{
		"[G]     [J] [N] [X] [Z]     [M] [V]",
		"[H] [B] [R] [S] [R] [T] [S] [R] [L]",
		" 1   2   3   4   5   6   7   8   9 ",
	})

	supplyArea.MoveSupply(2, 1, 3)

	if len(supplyArea.areas[0]) != 0 {
		t.Error("Did not remove supply from correct area")
	}

	if len(supplyArea.areas[2]) != 4 {
		t.Error("Did not add supply to correct area")
	}

	if supplyArea.areas[2][1] != "J" || supplyArea.areas[2][2] != "G" || supplyArea.areas[2][3] != "H" {
		t.Error("Did not add supply to correct area", supplyArea, "expected R J G H")
	}
}
