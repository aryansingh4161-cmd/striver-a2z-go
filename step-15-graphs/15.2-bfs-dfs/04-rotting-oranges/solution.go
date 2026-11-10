package rottingoranges

// OrangesRotting returns minutes until all oranges rot, or -1. Multi-source BFS. O(n*m).
func OrangesRotting(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	g := make([][]int, n)
	for i := range grid {
		g[i] = append([]int{}, grid[i]...)
	}
	var q [][2]int
	fresh := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			switch g[r][c] {
			case 2:
				q = append(q, [2]int{r, c})
			case 1:
				fresh++
			}
		}
	}
	minutes := 0
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 && fresh > 0 {
		minutes++
		for size := len(q); size > 0; size-- {
			cur := q[0]
			q = q[1:]
			for _, d := range dirs {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r >= 0 && c >= 0 && r < n && c < m && g[r][c] == 1 {
					g[r][c] = 2
					fresh--
					q = append(q, [2]int{r, c})
				}
			}
		}
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}
