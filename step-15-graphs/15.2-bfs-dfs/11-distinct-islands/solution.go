package distinctislands

import (
	"fmt"
	"strings"
)

// CountDistinct returns the number of island shapes (translation only). O(n*m).
func CountDistinct(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, m)
	}
	shapes := map[string]bool{}
	var sb strings.Builder
	var dfs func(r, c, r0, c0 int)
	dfs = func(r, c, r0, c0 int) {
		if r < 0 || c < 0 || r >= n || c >= m || grid[r][c] != 1 || seen[r][c] {
			return
		}
		seen[r][c] = true
		fmt.Fprintf(&sb, "%d,%d;", r-r0, c-c0)
		dfs(r+1, c, r0, c0)
		dfs(r-1, c, r0, c0)
		dfs(r, c+1, r0, c0)
		dfs(r, c-1, r0, c0)
	}
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 && !seen[r][c] {
				sb.Reset()
				dfs(r, c, r, c)
				shapes[sb.String()] = true
			}
		}
	}
	return len(shapes)
}
