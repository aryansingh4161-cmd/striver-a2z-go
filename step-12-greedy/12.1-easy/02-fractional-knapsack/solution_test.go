package fractionalknapsack

import (
	"math"
	"testing"
)

func TestMaxValue(t *testing.T) {
	tests := []struct {
		v, w []int
		cap  int
		want float64
	}{
		{[]int{60, 100, 120}, []int{10, 20, 30}, 50, 240},
		{[]int{60, 100}, []int{10, 20}, 50, 160},
		{[]int{500}, []int{30}, 10, 166.6666666667},
		{[]int{10}, []int{5}, 0, 0},
	}
	for _, tc := range tests {
		if got := MaxValue(tc.v, tc.w, tc.cap); math.Abs(got-tc.want) > 1e-6 {
			t.Errorf("MaxValue = %v, want %v", got, tc.want)
		}
	}
}
