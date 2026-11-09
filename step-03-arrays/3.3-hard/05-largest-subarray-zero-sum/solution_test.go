package largestsubarrayzerosum

import "testing"

func TestLongestZeroSum(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{15, -2, 2, -8, 1, 7, 10, 23}, 5},
		{[]int{1, 2, 3}, 0},
		{[]int{1, 0, 3}, 1},
		{[]int{1, -1, 1, -1}, 4},
	}
	for _, tc := range tests {
		if got := LongestZeroSum(tc.a); got != tc.want {
			t.Errorf("LongestZeroSum(%v) = %d", tc.a, got)
		}
	}
}
