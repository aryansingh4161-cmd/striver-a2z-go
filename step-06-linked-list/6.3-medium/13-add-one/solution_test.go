package addone

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestAddOne(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{1, 5, 9}, []int{1, 6, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{0}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 4}},
	}
	for _, tc := range tests {
		for name, fn := range map[string]func(*Node) *Node{"reverse": AddOneReverse, "rec": AddOneRecursive} {
			if got := list.ToSlice(fn(list.FromSlice(tc.a))); !slices.Equal(got, tc.want) {
				t.Errorf("%s(%v) = %v", name, tc.a, got)
			}
		}
	}
}
