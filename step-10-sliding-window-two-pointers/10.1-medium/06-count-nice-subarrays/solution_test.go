package countnicesubarrays

import "testing"

func TestNumberOfSubarrays(t *testing.T) {
	tests := []struct {
		a       []int
		k, want int
	}{{[]int{1, 1, 2, 1, 1}, 3, 2}, {[]int{2, 4, 6}, 1, 0}, {[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2, 16}}
	for _, tc := range tests {
		if got := NumberOfSubarrays(tc.a, tc.k); got != tc.want {
			t.Errorf("NumberOfSubarrays(%v,%d) = %d", tc.a, tc.k, got)
		}
	}
}
