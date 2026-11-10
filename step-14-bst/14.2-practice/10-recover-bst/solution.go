package recoverbst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// Recover swaps back the two misplaced values. O(n) time, O(h) space.
func Recover(root *Node) {
	var first, middle, last, prev *Node
	var inorder func(*Node)
	inorder = func(n *Node) {
		if n == nil {
			return
		}
		inorder(n.Left)
		if prev != nil && prev.Val > n.Val {
			if first == nil {
				first, middle = prev, n
			} else {
				last = n
			}
		}
		prev = n
		inorder(n.Right)
	}
	inorder(root)
	if first == nil {
		return
	}
	if last != nil {
		first.Val, last.Val = last.Val, first.Val
	} else {
		first.Val, middle.Val = middle.Val, first.Val
	}
}
