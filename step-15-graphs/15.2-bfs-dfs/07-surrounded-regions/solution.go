package surroundedregions

// Solve flips surrounded 'O' regions to 'X' in place. O(n*m).
func Solve(b [][]byte) {
	n := len(b)
	if n == 0 {
		return
	}
	m := len(b[0])
	safe := make([][]bool, n)
	for i := range safe {
		safe[i] = make([]bool, m)
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= n || c >= m || b[r][c] != 'O' || safe[r][c] {
			return
		}
		safe[r][c] = true
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	for r := 0; r < n; r++ {
		dfs(r, 0)
		dfs(r, m-1)
	}
	for c := 0; c < m; c++ {
		dfs(0, c)
		dfs(n-1, c)
	}
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if b[r][c] == 'O' && !safe[r][c] {
				b[r][c] = 'X'
			}
		}
	}
}
