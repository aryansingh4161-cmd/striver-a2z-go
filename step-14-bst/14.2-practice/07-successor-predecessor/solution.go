package successorpredecessor

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// Successor returns the smallest value > x, or -1. O(h).
func Successor(n *Node, x int) int {
	ans := -1
	for n != nil {
		if n.Val > x {
			ans = n.Val
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return ans
}

// Predecessor returns the largest value < x, or -1. O(h).
func Predecessor(n *Node, x int) int {
	ans := -1
	for n != nil {
		if n.Val < x {
			ans = n.Val
			n = n.Right
		} else {
			n = n.Left
		}
	}
	return ans
}
