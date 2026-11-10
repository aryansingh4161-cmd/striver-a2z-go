package checkbalanced

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// IsBalanced reports whether every node's subtrees differ in height by <= 1. O(n).
func IsBalanced(root *Node) bool {
	return height(root) != -1
}

// height returns the height, or -1 if the subtree is unbalanced.
func height(n *Node) int {
	if n == nil {
		return 0
	}
	l := height(n.Left)
	if l == -1 {
		return -1
	}
	r := height(n.Right)
	if r == -1 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	return 1 + max(l, r)
}
