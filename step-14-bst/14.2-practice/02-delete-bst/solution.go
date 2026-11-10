package deletebst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// Delete removes key and returns the new root. O(h).
func Delete(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	switch {
	case key < n.Val:
		n.Left = Delete(n.Left, key)
	case key > n.Val:
		n.Right = Delete(n.Right, key)
	default:
		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}
		succ := n.Right // smallest in the right subtree
		for succ.Left != nil {
			succ = succ.Left
		}
		n.Val = succ.Val
		n.Right = Delete(n.Right, succ.Val)
	}
	return n
}
