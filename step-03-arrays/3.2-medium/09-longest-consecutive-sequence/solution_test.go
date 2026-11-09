package longestconsecutivesequence

import "testing"

func TestLongest(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{}, 0},
		{[]int{1, 2, 0, 1}, 3},
		{[]int{-2, -1, 5}, 2},
	}
	for _, tc := range tests {
		if got := Longest(tc.a); got != tc.want {
			t.Errorf("Longest(%v) = %d", tc.a, got)
		}
		if got := LongestSort(tc.a); got != tc.want {
			t.Errorf("LongestSort(%v) = %d", tc.a, got)
		}
	}
}
