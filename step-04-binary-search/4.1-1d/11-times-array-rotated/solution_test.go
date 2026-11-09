package timesarrayrotated

import "testing"

func TestRotations(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 4},
		{[]int{3, 4, 5, 1, 2}, 3},
		{[]int{1, 2, 3}, 0},
		{[]int{2, 1}, 1},
	}
	for _, tc := range tests {
		if got := Rotations(tc.a); got != tc.want {
			t.Errorf("Rotations(%v) = %d", tc.a, got)
		}
	}
}
