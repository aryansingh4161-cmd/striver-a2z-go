package editdistance

// Recursive tries insert / delete / replace. Exponential.
func Recursive(a, b string) int {
	var f func(i, j int) int
	f = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		if a[i] == b[j] {
			return f(i-1, j-1)
		}
		return 1 + min(f(i, j-1), f(i-1, j), f(i-1, j-1))
	}
	return f(len(a)-1, len(b)-1)
}

// Memo caches the shifted f(i, j). O(n·m).
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
		if i == 0 {
			return j
		}
		if j == 0 {
			return i
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		if a[i-1] == b[j-1] {
			dp[i][j] = f(i-1, j-1)
		} else {
			dp[i][j] = 1 + min(f(i, j-1), f(i-1, j), f(i-1, j-1))
		}
		return dp[i][j]
	}
	return f(n, m)
}

// Optimal keeps two rows. O(n·m) time, O(m) space.
func Optimal(a, b string) int {
	m := len(b)
	prev, cur := make([]int, m+1), make([]int, m+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(a); i++ {
		cur[0] = i
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				cur[j] = prev[j-1]
			} else {
				cur[j] = 1 + min(cur[j-1], prev[j], prev[j-1])
			}
		}
		prev, cur = cur, prev
	}
	return prev[m]
}
