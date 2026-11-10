// Package bst has small helpers for Step 14 tests. Nodes are the same type as
// in Step 13 (tree.Node).
package bst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Build inserts values one by one into an empty BST and returns the root.
func Build(vals []int) *Node {
	var root *Node
	for _, v := range vals {
		root = insert(root, v)
	}
	return root
}

func insert(n *Node, v int) *Node {
	if n == nil {
		return &Node{Val: v}
	}
	if v < n.Val {
		n.Left = insert(n.Left, v)
	} else {
		n.Right = insert(n.Right, v)
	}
	return n
}

// Inorder returns the inorder values (sorted for a valid BST).
func Inorder(n *Node) []int {
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
	f(n)
	return out
}
