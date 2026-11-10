package printlcs

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.4-strings/01-lcs"

// PrintLCS returns one longest common subsequence of a and b.
func PrintLCS(a, b string) string {
	dp := lcs.Table(a, b)
	i, j := len(a), len(b)
	out := make([]byte, dp[i][j])
	k := len(out) - 1 // fill from the back
	for i > 0 && j > 0 {
		switch {
		case a[i-1] == b[j-1]:
			out[k] = a[i-1]
			k--
			i--
			j--
		case dp[i-1][j] >= dp[i][j-1]:
			i--
		default:
			j--
		}
	}
	return string(out)
}
