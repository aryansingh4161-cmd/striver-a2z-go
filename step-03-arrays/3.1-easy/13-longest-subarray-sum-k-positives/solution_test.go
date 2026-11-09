package longestsubarraysumkpositives

import "testing"

func TestLongest(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want int
	}{
		{[]int{1, 2, 3, 1, 1, 1, 1, 4, 2, 3}, 3, 3},
		{[]int{2, 0, 0, 3}, 3, 3},
		{[]int{1, 2, 3}, 10, 0},
		{[]int{0, 0, 0}, 0, 3},
	}
	for _, tc := range tests {
		if got := Longest(tc.a, tc.k); got != tc.want {
			t.Errorf("Longest(%v,%d) = %d, want %d", tc.a, tc.k, got, tc.want)
		}
		if got := LongestBrute(tc.a, tc.k); got != tc.want {
			t.Errorf("LongestBrute(%v,%d) = %d, want %d", tc.a, tc.k, got, tc.want)
		}
	}
}
