package postorderonestack

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestPostorder(t *testing.T) {
	cases := []struct{ lvl, want []int }{
		{[]int{1, 2, 3, 4, 5, tree.Null, 6}, []int{4, 5, 2, 6, 3, 1}},
		{[]int{1, tree.Null, 2, 3}, []int{3, 2, 1}},
		{[]int{1}, []int{1}},
	}
	for _, c := range cases {
		if got := Postorder(tree.FromLevel(c.lvl)); !slices.Equal(got, c.want) {
			t.Errorf("Postorder(%v) = %v", c.lvl, got)
		}
	}
}
