package childrensumproperty

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func holds(n *Node) bool {
	if n == nil || (n.Left == nil && n.Right == nil) {
		return true
	}
	s := 0
	if n.Left != nil {
		s += n.Left.Val
	}
	if n.Right != nil {
		s += n.Right.Val
	}
	return n.Val == s && holds(n.Left) && holds(n.Right)
}

func TestFix(t *testing.T) {
	N := tree.Null
	for _, lvl := range [][]int{{40, 10, 20, 2, 5, 30, 40}, {50, 7, 2, 3, 5, 1, 30}, {2, 35, 10, 2, 3, 5, 2}, {5, 1, N, 1}} {
		before := tree.ToLevel(tree.FromLevel(lvl))
		r := tree.FromLevel(lvl)
		Fix(r)
		if !holds(r) {
			t.Errorf("property fails for %v → %v", lvl, tree.ToLevel(r))
		}
		after := tree.ToLevel(r)
		for i := range before {
			if before[i] != N && after[i] < before[i] {
				t.Errorf("value decreased in %v", lvl)
			}
		}
	}
}
