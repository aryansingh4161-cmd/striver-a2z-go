package ceilbst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// Ceil returns the smallest value >= x, or -1. O(h).
func Ceil(n *Node, x int) int {
	ans := -1
	for n != nil {
		switch {
		case n.Val == x:
			return x
		case n.Val > x:
			ans = n.Val
			n = n.Left
		default:
			n = n.Right
		}
	}
	return ans
}
