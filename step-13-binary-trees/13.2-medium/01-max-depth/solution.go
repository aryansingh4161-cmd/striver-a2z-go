package maxdepth

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// MaxDepth returns the height of the tree (recursive). O(n).
func MaxDepth(n *Node) int {
	if n == nil {
		return 0
	}
	return 1 + max(MaxDepth(n.Left), MaxDepth(n.Right))
}

// MaxDepthBFS counts levels with a queue. O(n).
func MaxDepthBFS(root *Node) int {
	if root == nil {
		return 0
	}
	q, depth := []*Node{root}, 0
	for len(q) > 0 {
		depth++
		for size := len(q); size > 0; size-- {
			n := q[0]
			q = q[1:]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
	}
	return depth
}
