package movezerostoend

import (
	"slices"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{1, 0, 2, 3, 0, 4, 0, 1}, []int{1, 2, 3, 4, 1, 0, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1, 2}, []int{1, 2}},
	}
	for _, tc := range tests {
		a := append([]int{}, tc.a...)
		MoveZeros(a)
		if !slices.Equal(a, tc.want) {
			t.Errorf("MoveZeros(%v) = %v", tc.a, a)
		}
	}
}
