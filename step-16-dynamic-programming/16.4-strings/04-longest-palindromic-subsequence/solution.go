package longestpalindromicsubsequence

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.4-strings/01-lcs"

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// LPS = LCS(s, reverse(s)). O(n²) time, O(n) space.
func LPS(s string) int { return lcs.Optimal(s, reverse(s)) }

// Interval is the direct f(i, j) over s[i..j], bottom-up. O(n²).
func Interval(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1] // dp[i+1][i] is 0 when j == i+1
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
