package floorandceil

import "testing"

func TestFloorCeil(t *testing.T) {
	a := []int{3, 4, 4, 7, 8, 10}
	tests := []struct{ x, floor, ceil int }{{5, 4, 7}, {8, 8, 8}, {2, -1, 3}, {11, 10, -1}}
	for _, tc := range tests {
		if got := Floor(a, tc.x); got != tc.floor {
			t.Errorf("Floor(%d) = %d", tc.x, got)
		}
		if got := Ceil(a, tc.x); got != tc.ceil {
			t.Errorf("Ceil(%d) = %d", tc.x, got)
		}
	}
}
