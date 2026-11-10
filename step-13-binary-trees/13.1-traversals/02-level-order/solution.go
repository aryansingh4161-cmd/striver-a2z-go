package levelorder

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// LevelOrder returns node values grouped by level. Time O(n), space O(n).
func LevelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*Node{root}
	for len(q) > 0 {
		size := len(q)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			level = append(level, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
