package wildcardmatching

func allStars(p string, upto int) bool {
	for k := 0; k <= upto; k++ {
		if p[k] != '*' {
			return false
		}
	}
	return true
}

// Recursive tries both meanings of '*'. Exponential.
func Recursive(s, p string) bool {
	var f func(i, j int) bool // pattern p[0..i], string s[0..j]
	f = func(i, j int) bool {
		if i < 0 && j < 0 {
			return true
		}
		if i < 0 {
			return false
		}
		if j < 0 {
			return allStars(p, i)
		}
		if p[i] == s[j] || p[i] == '?' {
			return f(i-1, j-1)
		}
		if p[i] == '*' {
			return f(i-1, j) || f(i, j-1)
		}
		return false
	}
	return f(len(p)-1, len(s)-1)
}

// Tab is the shifted table: dp[i][j] = p[:i] matches s[:j]. O(n·m).
func Tab(s, p string) bool {
	n, m := len(p), len(s)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] && p[i-1] == '*'
		for j := 1; j <= m; j++ {
			switch {
			case p[i-1] == s[j-1] || p[i-1] == '?':
				dp[i][j] = dp[i-1][j-1]
			case p[i-1] == '*':
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}

// Optimal keeps two rows. O(m) space.
func Optimal(s, p string) bool {
	m := len(s)
	prev, cur := make([]bool, m+1), make([]bool, m+1)
	prev[0] = true
	for i := 1; i <= len(p); i++ {
		cur[0] = prev[0] && p[i-1] == '*'
		for j := 1; j <= m; j++ {
			switch {
			case p[i-1] == s[j-1] || p[i-1] == '?':
				cur[j] = prev[j-1]
			case p[i-1] == '*':
				cur[j] = prev[j] || cur[j-1]
			default:
				cur[j] = false
			}
		}
		prev, cur = cur, prev
	}
	return prev[m]
}
