package splitarraylargestsum

import "testing"

func TestSplitArray(t *testing.T) {
	tests := []struct {
		a       []int
		k, want int
	}{
		{[]int{7, 2, 5, 10, 8}, 2, 18},
		{[]int{1, 2, 3, 4, 5}, 2, 9},
		{[]int{1, 4, 4}, 3, 4},
	}
	for _, tc := range tests {
		if got := SplitArray(tc.a, tc.k); got != tc.want {
			t.Errorf("SplitArray(%v,%d) = %d", tc.a, tc.k, got)
		}
	}
}
