package minmaxbst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// Min returns the smallest value (root must be non-nil). O(h).
func Min(n *Node) int {
	for n.Left != nil {
		n = n.Left
	}
	return n.Val
}

// Max returns the largest value (root must be non-nil). O(h).
func Max(n *Node) int {
	for n.Right != nil {
		n = n.Right
	}
	return n.Val
}
