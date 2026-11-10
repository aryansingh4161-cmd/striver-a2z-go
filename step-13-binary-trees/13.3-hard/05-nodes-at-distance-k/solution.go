package nodesatdistancek

import (
	"sort"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

type Node = tree.Node

// DistanceK returns the values at distance k from target, sorted. O(n).
func DistanceK(root, target *Node, k int) []int {
	parent := map[*Node]*Node{}
	var mark func(n *Node)
	mark = func(n *Node) {
		for _, c := range []*Node{n.Left, n.Right} {
			if c != nil {
				parent[c] = n
				mark(c)
			}
		}
	}
	if root != nil {
		mark(root)
	}
	visited := map[*Node]bool{target: true}
	q := []*Node{target}
	for d := 0; d < k && len(q) > 0; d++ {
		for size := len(q); size > 0; size-- {
			n := q[0]
			q = q[1:]
			for _, nb := range []*Node{n.Left, n.Right, parent[n]} {
				if nb != nil && !visited[nb] {
					visited[nb] = true
					q = append(q, nb)
				}
			}
		}
	}
	res := []int{}
	for _, n := range q {
		res = append(res, n.Val)
	}
	sort.Ints(res)
	return res
}
