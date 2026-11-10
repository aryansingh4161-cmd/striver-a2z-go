package longestcommonsubstring

// Brute tries every pair of start positions. O(n·m·min(n,m)).
func Brute(a, b string) int {
	best := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			k := 0
			for i+k < len(a) && j+k < len(b) && a[i+k] == b[j+k] {
				k++
			}
			best = max(best, k)
		}
	}
	return best
}

// Tab: dp[i][j] = common run ending at a[i-1], b[j-1]. O(n·m).
func Tab(a, b string) int {
	n, m := len(a), len(b)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	best := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				best = max(best, dp[i][j])
			}
		}
	}
	return best
}

// Optimal keeps two rows. O(m) space.
func Optimal(a, b string) int {
	m := len(b)
	prev, cur := make([]int, m+1), make([]int, m+1)
	best := 0
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				cur[j] = 1 + prev[j-1]
				best = max(best, cur[j])
			} else {
				cur[j] = 0 // the run breaks
			}
		}
		prev, cur = cur, prev
	}
	return best
}
