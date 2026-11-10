package maxpathsum

import (
	"math"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

type Node = tree.Node

// MaxPathSum returns the largest path sum in the tree. O(n).
func MaxPathSum(root *Node) int {
	best := math.MinInt
	var gain func(*Node) int
	gain = func(n *Node) int {
		if n == nil {
			return 0
		}
		l := max(0, gain(n.Left))
		r := max(0, gain(n.Right))
		best = max(best, n.Val+l+r)
		return n.Val + max(l, r)
	}
	gain(root)
	return best
}
