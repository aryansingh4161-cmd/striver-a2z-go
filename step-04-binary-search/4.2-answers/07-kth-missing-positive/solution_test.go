package kthmissingpositive

import "testing"

func TestKthMissing(t *testing.T) {
	tests := []struct {
		a       []int
		k, want int
	}{
		{[]int{2, 3, 4, 7, 11}, 5, 9},
		{[]int{1, 2, 3, 4}, 2, 6},
		{[]int{5, 6, 7}, 3, 3},
		{[]int{}, 4, 4},
	}
	for _, tc := range tests {
		if got := KthMissing(tc.a, tc.k); got != tc.want {
			t.Errorf("KthMissing(%v,%d) = %d", tc.a, tc.k, got)
		}
		if got := KthMissingLinear(tc.a, tc.k); got != tc.want {
			t.Errorf("KthMissingLinear(%v,%d) = %d", tc.a, tc.k, got)
		}
	}
}
