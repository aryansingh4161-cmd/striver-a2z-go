package searchbst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// Search returns the node with value x, or nil. O(h), O(1) space.
func Search(n *Node, x int) *Node {
	for n != nil && n.Val != x {
		if x < n.Val {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return n
}
