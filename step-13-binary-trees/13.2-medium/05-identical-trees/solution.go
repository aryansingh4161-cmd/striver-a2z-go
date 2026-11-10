package identicaltrees

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// IsSame reports whether two trees have the same shape and values. O(n).
func IsSame(a, b *Node) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && IsSame(a.Left, b.Left) && IsSame(a.Right, b.Right)
}
