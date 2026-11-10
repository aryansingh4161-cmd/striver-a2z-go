package symmetrictree

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// IsSymmetric reports whether the tree mirrors itself. O(n).
func IsSymmetric(root *Node) bool {
	if root == nil {
		return true
	}
	return mirror(root.Left, root.Right)
}

func mirror(a, b *Node) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && mirror(a.Left, b.Right) && mirror(a.Right, b.Left)
}
