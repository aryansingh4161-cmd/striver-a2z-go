package recursivetraversals

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Preorder returns Root, Left, Right. O(n).
func Preorder(root *Node) []int {
	out := []int{}
	var f func(*Node)
	f = func(n *Node) {
		if n == nil {
			return
		}
		out = append(out, n.Val)
		f(n.Left)
		f(n.Right)
	}
	f(root)
	return out
}

// Inorder returns Left, Root, Right. O(n).
func Inorder(root *Node) []int {
	out := []int{}
	var f func(*Node)
	f = func(n *Node) {
		if n == nil {
			return
		}
		f(n.Left)
		out = append(out, n.Val)
		f(n.Right)
	}
	f(root)
	return out
}

// Postorder returns Left, Right, Root. O(n).
func Postorder(root *Node) []int {
	out := []int{}
	var f func(*Node)
	f = func(n *Node) {
		if n == nil {
			return
		}
		f(n.Left)
		f(n.Right)
		out = append(out, n.Val)
	}
	f(root)
	return out
}
