package palindromepartitioningii

func isPal(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// Recursive cuts off every palindromic prefix. Exponential.
func Recursive(s string) int {
	var f func(i int) int // fewest pieces for s[i:]
	f = func(i int) int {
		if i == len(s) {
			return 0
		}
		best := len(s)
		for j := i; j < len(s); j++ {
			if isPal(s, i, j) {
				best = min(best, 1+f(j+1))
			}
		}
		return best
	}
	return f(0) - 1
}

// Tab uses a precomputed palindrome table and f(i) from the end. O(n²).
func Tab(s string) int {
	n := len(s)
	pal := make([][]bool, n)
	for i := range pal {
		pal[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			pal[i][j] = s[i] == s[j] && (j-i < 2 || pal[i+1][j-1])
		}
	}
	dp := make([]int, n+1) // dp[i] = fewest pieces for s[i:]
	for i := n - 1; i >= 0; i-- {
		dp[i] = n
		for j := i; j < n; j++ {
			if pal[i][j] {
				dp[i] = min(dp[i], 1+dp[j+1])
			}
		}
	}
	return dp[0] - 1
}
