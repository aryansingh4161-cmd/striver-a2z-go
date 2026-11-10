package kthsmallestlargest

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// KthSmallest returns the k-th smallest value (1-based). O(h + k).
func KthSmallest(root *Node, k int) int {
	st := []*Node{}
	cur := root
	for {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
}

// KthLargest returns the k-th largest value using reverse inorder. O(h + k).
func KthLargest(root *Node, k int) int {
	st := []*Node{}
	cur := root
	for {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Right
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Left
	}
}
