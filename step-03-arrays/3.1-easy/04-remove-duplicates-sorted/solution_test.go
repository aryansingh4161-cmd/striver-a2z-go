package removeduplicatessorted

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{1, 1, 2, 2, 2, 3, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{5}, []int{5}},
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tc := range tests {
		a := append([]int{}, tc.a...)
		k := RemoveDuplicates(a)
		if !slices.Equal(a[:k], tc.want) {
			t.Errorf("RemoveDuplicates(%v) front = %v, want %v", tc.a, a[:k], tc.want)
		}
	}
}
