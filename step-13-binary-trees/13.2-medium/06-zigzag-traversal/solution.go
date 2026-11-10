package zigzagtraversal

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Zigzag returns level order with alternating direction. O(n).
func Zigzag(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*Node{root}
	leftToRight := true
	for len(q) > 0 {
		size := len(q)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			if leftToRight {
				level[i] = n.Val
			} else {
				level[size-1-i] = n.Val
			}
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		res = append(res, level)
		leftToRight = !leftToRight
	}
	return res
}
