package lcabst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// LCA returns the lowest common ancestor of values p and q. O(h), O(1).
func LCA(n *Node, p, q int) *Node {
	for n != nil {
		switch {
		case p < n.Val && q < n.Val:
			n = n.Left
		case p > n.Val && q > n.Val:
			n = n.Right
		default:
			return n
		}
	}
	return nil
}
