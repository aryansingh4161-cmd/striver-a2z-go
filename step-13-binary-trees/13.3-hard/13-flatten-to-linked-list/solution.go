package flattentolinkedlist

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// FlattenRecursive uses reverse preorder with a prev pointer. O(n), O(h).
func FlattenRecursive(root *Node) {
	var prev *Node
	var f func(*Node)
	f = func(n *Node) {
		if n == nil {
			return
		}
		f(n.Right)
		f(n.Left)
		n.Right = prev
		n.Left = nil
		prev = n
	}
	f(root)
}

// Flatten rewires in place with O(1) extra space.
func Flatten(root *Node) {
	for cur := root; cur != nil; cur = cur.Right {
		if cur.Left != nil {
			pred := cur.Left
			for pred.Right != nil {
				pred = pred.Right
			}
			pred.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
	}
}
