package nextgreaterelementii

import (
	"slices"
	"testing"
)

func TestNGECircular(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{1, 2, 1}, []int{2, -1, 2}},
		{[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
		{[]int{5}, []int{-1}},
		{[]int{3, 3, 3}, []int{-1, -1, -1}},
	}
	for _, tc := range tests {
		if got := NGECircular(tc.a); !slices.Equal(got, tc.want) {
			t.Errorf("NGECircular(%v) = %v", tc.a, got)
		}
	}
}
