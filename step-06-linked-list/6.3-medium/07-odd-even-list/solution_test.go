package oddevenlist

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestOddEven(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
		{[]int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{}, []int{}},
	}
	for _, tc := range tests {
		if got := list.ToSlice(OddEven(list.FromSlice(tc.a))); !slices.Equal(got, tc.want) {
			t.Errorf("OddEven(%v) = %v", tc.a, got)
		}
	}
}
