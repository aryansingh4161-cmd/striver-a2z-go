package largestbstinbt

import (
	"math"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

type Node = tree.Node

// LargestBST returns the node count of the largest BST subtree. O(n).
func LargestBST(root *Node) int {
	_, _, size := info(root)
	return size
}

// info returns (min, max, size of largest BST inside).
func info(n *Node) (int, int, int) {
	if n == nil {
		return math.MaxInt, math.MinInt, 0
	}
	lmin, lmax, lsize := info(n.Left)
	rmin, rmax, rsize := info(n.Right)
	if lmax < n.Val && n.Val < rmin {
		return min(lmin, n.Val), max(rmax, n.Val), lsize + rsize + 1
	}
	return math.MinInt, math.MaxInt, max(lsize, rsize)
}
