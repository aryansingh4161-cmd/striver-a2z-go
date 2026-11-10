package topview

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// TopView returns the topmost node of each column, left to right. O(n).
func TopView(root *Node) []int {
	if root == nil {
		return []int{}
	}
	type item struct {
		n   *Node
		col int
	}
	first := map[int]int{}
	lo, hi := 0, 0
	q := []item{{root, 0}}
	for len(q) > 0 {
		it := q[0]
		q = q[1:]
		if _, ok := first[it.col]; !ok {
			first[it.col] = it.n.Val
		}
		lo, hi = min(lo, it.col), max(hi, it.col)
		if it.n.Left != nil {
			q = append(q, item{it.n.Left, it.col - 1})
		}
		if it.n.Right != nil {
			q = append(q, item{it.n.Right, it.col + 1})
		}
	}
	res := make([]int, 0, hi-lo+1)
	for c := lo; c <= hi; c++ {
		res = append(res, first[c])
	}
	return res
}
