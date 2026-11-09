package nextpermutation

import (
	"slices"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{2, 1, 5, 4, 3, 0, 0}, []int{2, 3, 0, 0, 1, 4, 5}},
		{[]int{1}, []int{1}},
	}
	for _, tc := range tests {
		a := append([]int{}, tc.a...)
		NextPermutation(a)
		if !slices.Equal(a, tc.want) {
			t.Errorf("NextPermutation(%v) = %v, want %v", tc.a, a, tc.want)
		}
	}
}
