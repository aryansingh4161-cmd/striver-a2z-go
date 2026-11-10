package distinctsubsequences

// Recursive: on a match, use s[i] or skip it. Exponential.
func Recursive(s, t string) int {
	var f func(i, j int) int
	f = func(i, j int) int {
		if j < 0 {
			return 1
		}
		if i < 0 {
			return 0
		}
		if s[i] == t[j] {
			return f(i-1, j-1) + f(i-1, j)
		}
		return f(i-1, j)
	}
	return f(len(s)-1, len(t)-1)
}

// Tab is the shifted table. O(n·m) time and space.
func Tab(s, t string) int {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1 // empty t is always formed once
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}

// Optimal uses one row with j going downward. O(m) space.
func Optimal(s, t string) int {
	m := len(t)
	dp := make([]int, m+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		for j := m; j >= 1; j-- {
			if s[i] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[m]
}
