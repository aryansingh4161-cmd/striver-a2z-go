package validatebst

import (
	"math"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

type Node = bst.Node

// IsValid checks the BST property with value ranges. O(n).
func IsValid(root *Node) bool {
	var f func(n *Node, low, high int) bool
	f = func(n *Node, low, high int) bool {
		if n == nil {
			return true
		}
		if n.Val <= low || n.Val >= high {
			return false
		}
		return f(n.Left, low, n.Val) && f(n.Right, n.Val, high)
	}
	return f(root, math.MinInt, math.MaxInt)
}
