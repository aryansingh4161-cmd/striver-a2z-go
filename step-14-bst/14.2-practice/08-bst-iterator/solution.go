package bstiterator

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// BSTIterator walks a BST in ascending order with O(h) memory.
type BSTIterator struct{ st []*Node }

// Constructor creates an iterator positioned before the smallest value.
func Constructor(root *Node) BSTIterator {
	it := BSTIterator{}
	it.pushLeft(root)
	return it
}

func (it *BSTIterator) pushLeft(n *Node) {
	for n != nil {
		it.st = append(it.st, n)
		n = n.Left
	}
}

// Next returns the next smallest value. Amortised O(1).
func (it *BSTIterator) Next() int {
	n := it.st[len(it.st)-1]
	it.st = it.st[:len(it.st)-1]
	it.pushLeft(n.Right)
	return n.Val
}

// HasNext reports whether more values remain.
func (it *BSTIterator) HasNext() bool { return len(it.st) > 0 }
