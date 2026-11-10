package boundarytraversal

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

func isLeaf(n *Node) bool { return n.Left == nil && n.Right == nil }

// Boundary returns the anticlockwise boundary starting from the root. O(n).
func Boundary(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	if !isLeaf(root) {
		res = append(res, root.Val)
	}
	// left boundary
	for cur := root.Left; cur != nil; {
		if !isLeaf(cur) {
			res = append(res, cur.Val)
		}
		if cur.Left != nil {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	// leaves
	var leaves func(*Node)
	leaves = func(n *Node) {
		if n == nil {
			return
		}
		if isLeaf(n) {
			res = append(res, n.Val)
			return
		}
		leaves(n.Left)
		leaves(n.Right)
	}
	leaves(root)
	// right boundary, reversed
	var right []int
	for cur := root.Right; cur != nil; {
		if !isLeaf(cur) {
			right = append(right, cur.Val)
		}
		if cur.Right != nil {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	for i := len(right) - 1; i >= 0; i-- {
		res = append(res, right[i])
	}
	return res
}
