package insertbst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// Insert adds x as a new leaf and returns the root. O(h), O(1) space.
func Insert(root *Node, x int) *Node {
	n := &Node{Val: x}
	if root == nil {
		return n
	}
	cur := root
	for {
		if x < cur.Val {
			if cur.Left == nil {
				cur.Left = n
				return root
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = n
				return root
			}
			cur = cur.Right
		}
	}
}
