package diameter

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Diameter returns the longest path length in edges. O(n).
func Diameter(root *Node) int {
	best := 0
	var h func(*Node) int
	h = func(n *Node) int {
		if n == nil {
			return 0
		}
		l, r := h(n.Left), h(n.Right)
		best = max(best, l+r)
		return 1 + max(l, r)
	}
	h(root)
	return best
}
