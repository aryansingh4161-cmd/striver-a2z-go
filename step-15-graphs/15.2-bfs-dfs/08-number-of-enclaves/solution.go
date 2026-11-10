package numberofenclaves

// NumEnclaves counts land cells that cannot reach the border. Multi-source BFS. O(n*m).
func NumEnclaves(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, m)
	}
	var q [][2]int
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if (r == 0 || c == 0 || r == n-1 || c == m-1) && grid[r][c] == 1 {
				seen[r][c] = true
				q = append(q, [2]int{r, c})
			}
		}
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, d := range dirs {
			r, c := cur[0]+d[0], cur[1]+d[1]
			if r >= 0 && c >= 0 && r < n && c < m && grid[r][c] == 1 && !seen[r][c] {
				seen[r][c] = true
				q = append(q, [2]int{r, c})
			}
		}
	}
	count := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == 1 && !seen[r][c] {
				count++
			}
		}
	}
	return count
}
