package distinctislands

import "testing"

func TestCountDistinct(t *testing.T) {
	if CountDistinct([][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}) != 1 {
		t.Error("same squares")
	}
	if CountDistinct([][]int{{1, 1, 0, 1, 1}, {1, 0, 0, 0, 0}, {0, 0, 0, 1, 1}, {1, 1, 0, 1, 0}}) != 2 { // two L-shapes + two 1x2 bars
		t.Error("two shapes")
	}
}
