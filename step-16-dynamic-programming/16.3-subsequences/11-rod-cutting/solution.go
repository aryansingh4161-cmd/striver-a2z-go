package rodcutting

// Recursive: for each piece length, skip it or cut it and stay. Exponential.
func Recursive(price []int, n int) int {
	var f func(i, L int) int
	f = func(i, L int) int {
		if i == 0 {
			return L * price[0]
		}
		best := f(i-1, L)
		if i+1 <= L {
			best = max(best, price[i]+f(i, L-(i+1)))
		}
		return best
	}
	return f(n-1, n)
}

// Memo caches f(i, L). O(n²).
func Memo(price []int, n int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(i, L int) int
	f = func(i, L int) int {
		if i == 0 {
			return L * price[0]
		}
		if dp[i][L] != -1 {
			return dp[i][L]
		}
		best := f(i-1, L)
		if i+1 <= L {
			best = max(best, price[i]+f(i, L-(i+1)))
		}
		dp[i][L] = best
		return best
	}
	return f(n-1, n)
}

// Optimal is the one-row unbounded knapsack. O(n²) time, O(n) space.
func Optimal(price []int, n int) int {
	dp := make([]int, n+1)
	for length := 1; length <= n; length++ {
		for L := length; L <= n; L++ {
			dp[L] = max(dp[L], price[length-1]+dp[L-length])
		}
	}
	return dp[n]
}
