package longestsubarraysumk

import "testing"

func TestLongest(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want int
	}{
		{[]int{1, -1, 5, -2, 3}, 3, 4},
		{[]int{-2, -1, 2, 1}, 1, 2},
		{[]int{2, 0, 0, 3}, 3, 3},
		{[]int{1, 2, 3}, 7, 0},
		{[]int{-1, 1, -1, 1}, 0, 4},
	}
	for _, tc := range tests {
		if got := Longest(tc.a, tc.k); got != tc.want {
			t.Errorf("Longest(%v,%d) = %d, want %d", tc.a, tc.k, got, tc.want)
		}
	}
}
