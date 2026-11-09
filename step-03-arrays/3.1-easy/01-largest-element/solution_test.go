package largestelement

import "testing"

func TestLargest(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{{[]int{3, 2, 1, 5, 2}, 5}, {[]int{-5, -2, -9}, -2}, {[]int{7}, 7}, {[]int{1, 1, 1}, 1}}
	for _, tc := range tests {
		if got := Largest(tc.a); got != tc.want {
			t.Errorf("Largest(%v) = %d", tc.a, got)
		}
		if got := LargestSort(tc.a); got != tc.want {
			t.Errorf("LargestSort(%v) = %d", tc.a, got)
		}
	}
}
