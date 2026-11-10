// Package tree holds the binary tree node and test helpers shared by Step 13.
package tree

import "math"

// Null marks a missing node in level-order input.
const Null = math.MinInt

// Node is a binary tree node.
type Node struct {
	Val         int
	Left, Right *Node
}

// FromLevel builds a tree from level-order values (Null = no node).
func FromLevel(vals []int) *Node {
	if len(vals) == 0 || vals[0] == Null {
		return nil
	}
	root := &Node{Val: vals[0]}
	queue := []*Node{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		n := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != Null {
			n.Left = &Node{Val: vals[i]}
			queue = append(queue, n.Left)
		}
		i++
		if i < len(vals) && vals[i] != Null {
			n.Right = &Node{Val: vals[i]}
			queue = append(queue, n.Right)
		}
		i++
	}
	return root
}

// ToLevel writes the tree in level order with Null for gaps (trailing Nulls removed).
func ToLevel(root *Node) []int {
	out := []int{}
	queue := []*Node{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n == nil {
			out = append(out, Null)
			continue
		}
		out = append(out, n.Val)
		queue = append(queue, n.Left, n.Right)
	}
	for len(out) > 0 && out[len(out)-1] == Null {
		out = out[:len(out)-1]
	}
	return out
}

// Find returns the first node with value v (preorder), or nil.
func Find(root *Node, v int) *Node {
	if root == nil || root.Val == v {
		return root
	}
	if l := Find(root.Left, v); l != nil {
		return l
	}
	return Find(root.Right, v)
}
