package uniquepathsii

// Recursive counts paths avoiding cells equal to 1. Exponential.
func Recursive(g [][]int) int {
	var f func(i, j int) int
	f = func(i, j int) int {
		if i < 0 || j < 0 || g[i][j] == 1 {
			return 0
		}
		if i == 0 && j == 0 {
			return 1
		}
		return f(i-1, j) + f(i, j-1)
	}
	return f(len(g)-1, len(g[0])-1)
}

// Memo caches f(i, j). O(m·n).
func Memo(g [][]int) int {
	m, n := len(g), len(g[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(i, j int) int
	f = func(i, j int) int {
		if i < 0 || j < 0 || g[i][j] == 1 {
			return 0
		}
		if i == 0 && j == 0 {
			return 1
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		dp[i][j] = f(i-1, j) + f(i, j-1)
		return dp[i][j]
	}
	return f(m-1, n-1)
}

// Tab is bottom-up. O(m·n) time and space.
func Tab(g [][]int) int {
	m, n := len(g), len(g[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			switch {
			case g[i][j] == 1:
				dp[i][j] = 0
			case i == 0 && j == 0:
				dp[i][j] = 1
			default:
				if i > 0 {
					dp[i][j] += dp[i-1][j]
				}
				if j > 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

// Optimal uses one row. O(m·n) time, O(n) space.
func Optimal(g [][]int) int {
	n := len(g[0])
	row := make([]int, n)
	for i := range g {
		for j := 0; j < n; j++ {
			switch {
			case g[i][j] == 1:
				row[j] = 0
			case i == 0 && j == 0:
				row[j] = 1
			case j > 0:
				row[j] += row[j-1]
			}
		}
	}
	return row[n-1]
}
