package floorbst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// Floor returns the largest value <= x, or -1. O(h).
func Floor(n *Node, x int) int {
	ans := -1
	for n != nil {
		switch {
		case n.Val == x:
			return x
		case n.Val < x:
			ans = n.Val
			n = n.Right
		default:
			n = n.Left
		}
	}
	return ans
}
