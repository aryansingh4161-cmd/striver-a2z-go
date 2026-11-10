package uniquepaths

// Recursive: f(i,j) = f(i-1,j) + f(i,j-1). Exponential.
func Recursive(m, n int) int {
	var f func(i, j int) int
	f = func(i, j int) int {
		if i == 0 && j == 0 {
			return 1
		}
		if i < 0 || j < 0 {
			return 0
		}
		return f(i-1, j) + f(i, j-1)
	}
	return f(m-1, n-1)
}

// Memo caches f(i, j). O(m·n).
func Memo(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(i, j int) int
	f = func(i, j int) int {
		if i == 0 && j == 0 {
			return 1
		}
		if i < 0 || j < 0 {
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		dp[i][j] = f(i-1, j) + f(i, j-1)
		return dp[i][j]
	}
	return f(m-1, n-1)
}

// Tab fills the full table. O(m·n) time and space.
func Tab(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// Optimal keeps one row. O(m·n) time, O(n) space.
func Optimal(m, n int) int {
	row := make([]int, n)
	for j := range row {
		row[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			row[j] += row[j-1] // row[j] was "above", row[j-1] is "left"
		}
	}
	return row[n-1]
}

// Combinatorics returns C(m+n-2, m-1). O(min(m, n)).
func Combinatorics(m, n int) int {
	if m > n {
		m, n = n, m
	}
	res := 1
	for i := 1; i < m; i++ {
		res = res * (n - 1 + i) / i // stays an integer at every step
	}
	return res
}
