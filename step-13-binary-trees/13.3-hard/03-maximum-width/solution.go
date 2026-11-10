package maximumwidth

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// WidthOfBinaryTree returns the maximum level width (nulls between count). O(n).
func WidthOfBinaryTree(root *Node) int {
	if root == nil {
		return 0
	}
	type item struct {
		n   *Node
		idx int
	}
	q := []item{{root, 0}}
	best := 0
	for len(q) > 0 {
		size := len(q)
		first := q[0].idx
		last := q[size-1].idx
		best = max(best, last-first+1)
		for i := 0; i < size; i++ {
			it := q[0]
			q = q[1:]
			id := it.idx - first // normalise to avoid overflow
			if it.n.Left != nil {
				q = append(q, item{it.n.Left, 2*id + 1})
			}
			if it.n.Right != nil {
				q = append(q, item{it.n.Right, 2*id + 2})
			}
		}
	}
	return best
}
