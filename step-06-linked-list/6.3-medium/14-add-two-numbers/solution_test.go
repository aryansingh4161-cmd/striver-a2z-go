package addtwonumbers

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestAddTwo(t *testing.T) {
	tests := []struct{ a, b, want []int }{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{[]int{5}, []int{5}, []int{0, 1}},
	}
	for _, tc := range tests {
		if got := list.ToSlice(AddTwo(list.FromSlice(tc.a), list.FromSlice(tc.b))); !slices.Equal(got, tc.want) {
			t.Errorf("AddTwo(%v,%v) = %v", tc.a, tc.b, got)
		}
	}
}
