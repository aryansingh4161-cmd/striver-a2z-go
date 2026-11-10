package burntree

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// BurnTime returns the seconds needed for fire from start to burn the whole tree. O(n).
func BurnTime(root, start *Node) int {
	parent := map[*Node]*Node{}
	var mark func(*Node)
	mark = func(n *Node) {
		for _, c := range []*Node{n.Left, n.Right} {
			if c != nil {
				parent[c] = n
				mark(c)
			}
		}
	}
	mark(root)
	visited := map[*Node]bool{start: true}
	q := []*Node{start}
	time := -1
	for len(q) > 0 {
		time++
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
	return time
}
