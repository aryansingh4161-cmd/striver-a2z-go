package replacebyrank

import (
	"slices"
	"testing"
)

func TestArrayRankTransform(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{20, 15, 26, 2, 98, 6}, []int{4, 3, 5, 1, 6, 2}},
		{[]int{40, 10, 20, 40}, []int{3, 1, 2, 3}},
		{[]int{100, 100, 100}, []int{1, 1, 1}},
		{[]int{}, []int{}},
	}
	for _, tc := range tests {
		if got := ArrayRankTransform(tc.a); !slices.Equal(got, tc.want) {
			t.Errorf("ArrayRankTransform(%v) = %v", tc.a, got)
		}
	}
}
