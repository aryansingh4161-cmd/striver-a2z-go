package reverselist

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestReverse(t *testing.T) {
	for name, fn := range map[string]func(*Node) *Node{"stack": ReverseStack, "iter": Reverse, "rec": ReverseRecursive} {
		for _, a := range [][]int{{1, 2, 3, 4}, {1}, {}, {5, 6}} {
			want := append([]int{}, a...)
			slices.Reverse(want)
			if got := list.ToSlice(fn(list.FromSlice(a))); !slices.Equal(got, want) {
				t.Errorf("%s(%v) = %v", name, a, got)
			}
		}
	}
}
