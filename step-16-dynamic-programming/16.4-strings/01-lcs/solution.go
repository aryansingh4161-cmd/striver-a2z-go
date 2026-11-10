// Package lcs has the longest common subsequence, reused by later problems.
package lcs

// Recursive compares last characters. Exponential.
func Recursive(a, b string) int {
	var f func(i, j int) int
	f = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if a[i] == b[j] {
			return 1 + f(i-1, j-1)
		}
		return max(f(i-1, j), f(i, j-1))
	}
	return f(len(a)-1, len(b)-1)
}

// Memo caches f(i, j) using shifted indexes (i = prefix length). O(n·m).
func Memo(a, b string) int {
	n, m := len(a), len(b)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(i, j int) int
	f = func(i, j int) int {
		if i == 0 || j == 0 {
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		if a[i-1] == b[j-1] {
			dp[i][j] = 1 + f(i-1, j-1)
		} else {
			dp[i][j] = max(f(i-1, j), f(i, j-1))
		}
		return dp[i][j]
	}
	return f(n, m)
}

// Table returns the full (n+1)×(m+1) LCS table. Used by Print LCS and SCS.
func Table(a, b string) [][]int {
	n, m := len(a), len(b)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp
}

// Tab is the bottom-up length. O(n·m) time and space.
func Tab(a, b string) int { return Table(a, b)[len(a)][len(b)] }

// Optimal keeps two rows. O(n·m) time, O(m) space.
func Optimal(a, b string) int {
	m := len(b)
	prev, cur := make([]int, m+1), make([]int, m+1)
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				cur[j] = 1 + prev[j-1]
			} else {
				cur[j] = max(prev[j], cur[j-1])
			}
		}
		prev, cur = cur, prev
	}
	return prev[m]
}
