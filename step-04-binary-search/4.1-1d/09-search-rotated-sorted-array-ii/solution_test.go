package searchrotatedsortedarrayii

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		a      []int
		target int
		want   bool
	}{
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{3, 1, 2, 3, 3, 3, 3}, 2, true},
		{[]int{1, 0, 1, 1, 1}, 0, true},
		{[]int{1, 1, 1, 1}, 2, false},
	}
	for _, tc := range tests {
		if got := Search(tc.a, tc.target); got != tc.want {
			t.Errorf("Search(%v,%d) = %v", tc.a, tc.target, got)
		}
	}
}
