package shippackages

import "testing"

func TestShip(t *testing.T) {
	tests := []struct {
		w          []int
		days, want int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
		{[]int{1, 2, 3, 1, 1}, 4, 3},
	}
	for _, tc := range tests {
		if got := ShipWithinDays(tc.w, tc.days); got != tc.want {
			t.Errorf("ShipWithinDays(%v,%d) = %d", tc.w, tc.days, got)
		}
	}
}
