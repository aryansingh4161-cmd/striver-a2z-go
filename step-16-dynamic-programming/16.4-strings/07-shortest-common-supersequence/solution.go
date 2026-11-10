package shortestcommonsupersequence

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-16-dynamic-programming/16.4-strings/01-lcs"

// SCS returns one shortest common supersequence of a and b. O(n·m).
func SCS(a, b string) string {
	dp := lcs.Table(a, b)
	i, j := len(a), len(b)
	out := make([]byte, 0, len(a)+len(b)-dp[i][j])
	for i > 0 && j > 0 {
		switch {
		case a[i-1] == b[j-1]:
			out = append(out, a[i-1])
			i--
			j--
		case dp[i-1][j] > dp[i][j-1]:
			out = append(out, a[i-1])
			i--
		default:
			out = append(out, b[j-1])
			j--
		}
	}
	for ; i > 0; i-- {
		out = append(out, a[i-1])
	}
	for ; j > 0; j-- {
		out = append(out, b[j-1])
	}
	for l, r := 0, len(out)-1; l < r; l, r = l+1, r-1 {
		out[l], out[r] = out[r], out[l]
	}
	return string(out)
}
