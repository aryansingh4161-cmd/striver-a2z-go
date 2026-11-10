package numberofislands

// NumIslands counts 4-directionally connected groups of '1'. O(n*m).
// The grid is not modified.
func NumIslands(grid [][]byte) int {
	return count(grid, [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}})
}

// NumIslands8 counts islands where diagonal neighbours also connect.
func NumIslands8(grid [][]byte) int {
	return count(grid, [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}})
}

func count(grid [][]byte, dirs [][2]int) int {
	if len(grid) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, m)
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= n || c >= m || grid[r][c] != '1' || seen[r][c] {
			return
		}
		seen[r][c] = true
		for _, d := range dirs {
			dfs(r+d[0], c+d[1])
		}
	}
	islands := 0
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if grid[r][c] == '1' && !seen[r][c] {
				islands++
				dfs(r, c)
			}
		}
	}
	return islands
}
