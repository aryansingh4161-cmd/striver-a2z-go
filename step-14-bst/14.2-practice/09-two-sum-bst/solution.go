package twosumbst

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"

type Node = bst.Node

// iter walks the BST ascending (reverse=false) or descending (reverse=true).
type iter struct {
	st      []*Node
	reverse bool
}

func newIter(root *Node, reverse bool) *iter {
	it := &iter{reverse: reverse}
	it.push(root)
	return it
}

func (it *iter) push(n *Node) {
	for n != nil {
		it.st = append(it.st, n)
		if it.reverse {
			n = n.Right
		} else {
			n = n.Left
		}
	}
}

func (it *iter) next() int {
	n := it.st[len(it.st)-1]
	it.st = it.st[:len(it.st)-1]
	if it.reverse {
		it.push(n.Left)
	} else {
		it.push(n.Right)
	}
	return n.Val
}

// FindTarget reports whether two distinct nodes sum to k. O(n) time, O(h) space.
func FindTarget(root *Node, k int) bool {
	if root == nil {
		return false
	}
	lo, hi := newIter(root, false), newIter(root, true)
	i, j := lo.next(), hi.next()
	for i < j {
		switch s := i + j; {
		case s == k:
			return true
		case s < k:
			i = lo.next()
		default:
			j = hi.next()
		}
	}
	return false
}
