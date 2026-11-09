package kokoeatingbananas

import "testing"

func TestMinSpeed(t *testing.T) {
	tests := []struct {
		piles   []int
		h, want int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
		{[]int{1}, 1, 1},
	}
	for _, tc := range tests {
		if got := MinSpeed(tc.piles, tc.h); got != tc.want {
			t.Errorf("MinSpeed(%v,%d) = %d", tc.piles, tc.h, got)
		}
	}
}
