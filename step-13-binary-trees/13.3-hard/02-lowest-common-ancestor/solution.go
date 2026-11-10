package lowestcommonancestor

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// LCA returns the lowest common ancestor of p and q (both in the tree). O(n).
func LCA(n, p, q *Node) *Node {
	if n == nil || n == p || n == q {
		return n
	}
	left, right := LCA(n.Left, p, q), LCA(n.Right, p, q)
	if left != nil && right != nil {
		return n
	}
	if left != nil {
		return left
	}
	return right
}
