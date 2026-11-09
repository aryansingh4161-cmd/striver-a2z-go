package spiralmatrix

import (
	"slices"
	"testing"
)

func TestSpiral(t *testing.T) {
	tests := []struct {
		m    [][]int
		want []int
	}{
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
	}
	for _, tc := range tests {
		if got := Spiral(tc.m); !slices.Equal(got, tc.want) {
			t.Errorf("Spiral(%v) = %v", tc.m, got)
		}
	}
}
