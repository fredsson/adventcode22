package main

import "testing"

func Test_CollapseSlices_ShouldCollapseSingleSlice(t *testing.T) {
	slices := []OccupiedSlice{{1, 10}}

	collapsed := CollapseSlices(slices)

	if len(collapsed) != 1 {
		t.Error("incorrect length from", slices, " result:", collapsed)
	}

	if collapsed[0].start != 1 || collapsed[0].end != 10 {
		t.Error("incorrect start/end", collapsed)
	}
}

func Test_CollapseSlices_ShouldCollapseToExtendedSlice(t *testing.T) {
	slices := []OccupiedSlice{{1, 10}, {5, 15}, {-10, 5}}

	collapsed := CollapseSlices(slices)

	if len(collapsed) != 1 {
		t.Error("incorrect length from", slices, " result:", collapsed)
	}

	if collapsed[0].start != -10 || collapsed[0].end != 15 {
		t.Error("incorrect start/end", collapsed)
	}
}

func Test_CollapseSlices_ShouldCollapseFullyContainedSlice(t *testing.T) {
	slices := []OccupiedSlice{{1, 10}, {5, 6}}

	collapsed := CollapseSlices(slices)

	if len(collapsed) != 1 {
		t.Error("incorrect length from", slices, " result:", collapsed)
	}

	if collapsed[0].start != 1 || collapsed[0].end != 10 {
		t.Error("incorrect start/end", collapsed)
	}
}

func Test_CollapseSlices_ShouldNotCollapseSliceToRight(t *testing.T) {
	slices := []OccupiedSlice{{1, 10}, {11, 22}}

	collapsed := CollapseSlices(slices)

	if len(collapsed) != 2 {
		t.Error("incorrect length from", slices, " result:", collapsed)
	}
}

func Test_CollapseSlices_ShouldNotCollapseSliceToLeft(t *testing.T) {
	slices := []OccupiedSlice{{1, 10}, {-10, -6}}

	collapsed := CollapseSlices(slices)

	if len(collapsed) != 2 {
		t.Error("incorrect length from", slices, " result:", collapsed)
	}
}
