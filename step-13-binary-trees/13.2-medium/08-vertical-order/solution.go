package verticalorder

import (
	"sort"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

type Node = tree.Node

// VerticalOrder groups nodes by column (sorted by row, then value). O(n log n).
func VerticalOrder(root *Node) [][]int {
	type entry struct{ col, row, val int }
	var all []entry
	var dfs func(n *Node, row, col int)
	dfs = func(n *Node, row, col int) {
		if n == nil {
			return
		}
		all = append(all, entry{col, row, n.Val})
		dfs(n.Left, row+1, col-1)
		dfs(n.Right, row+1, col+1)
	}
	dfs(root, 0, 0)
	sort.Slice(all, func(i, j int) bool {
		a, b := all[i], all[j]
		if a.col != b.col {
			return a.col < b.col
		}
		if a.row != b.row {
			return a.row < b.row
		}
		return a.val < b.val
	})
	res := [][]int{}
	for i, e := range all {
		if i == 0 || e.col != all[i-1].col {
			res = append(res, []int{})
		}
		res[len(res)-1] = append(res[len(res)-1], e.val)
	}
	return res
}
