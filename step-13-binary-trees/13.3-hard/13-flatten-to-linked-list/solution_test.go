package flattentolinkedlist

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func chain(n *Node) []int {
	out := []int{}
	for ; n != nil; n = n.Right {
		if n.Left != nil {
			return nil
		}
		out = append(out, n.Val)
	}
	return out
}

func TestFlatten(t *testing.T) {
	N := tree.Null
	for name, fn := range map[string]func(*Node){"rec": FlattenRecursive, "iter": Flatten} {
		r := tree.FromLevel([]int{1, 2, 5, 3, 4, N, 6})
		fn(r)
		if got := chain(r); !slices.Equal(got, []int{1, 2, 3, 4, 5, 6}) {
			t.Errorf("%s: %v", name, got)
		}
		fn(nil)
	}
}
