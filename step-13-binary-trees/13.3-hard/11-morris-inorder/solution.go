package morrisinorder

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Inorder returns the inorder traversal using Morris threading. O(n), O(1) extra.
func Inorder(root *Node) []int {
	out := []int{}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			out = append(out, cur.Val)
			cur = cur.Right
			continue
		}
		pred := cur.Left
		for pred.Right != nil && pred.Right != cur {
			pred = pred.Right
		}
		if pred.Right == nil {
			pred.Right = cur // create thread
			cur = cur.Left
		} else {
			pred.Right = nil // remove thread
			out = append(out, cur.Val)
			cur = cur.Right
		}
	}
	return out
}
