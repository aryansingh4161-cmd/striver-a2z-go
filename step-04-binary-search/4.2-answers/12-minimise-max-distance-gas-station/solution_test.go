package gasstation

import (
	"math"
	"testing"
)

func TestMinMaxDistance(t *testing.T) {
	tests := []struct {
		s    []int
		k    int
		want float64
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 0.5},
		{[]int{1, 2, 3, 4, 5}, 4, 0.5},
		{[]int{1, 7}, 2, 2.0},
		{[]int{1, 13, 17, 23}, 5, 3.0},
	}
	for _, tc := range tests {
		if got := MinMaxDistance(tc.s, tc.k); math.Abs(got-tc.want) > 1e-6 {
			t.Errorf("MinMaxDistance(%v,%d) = %v, want %v", tc.s, tc.k, got, tc.want)
		}
	}
}
