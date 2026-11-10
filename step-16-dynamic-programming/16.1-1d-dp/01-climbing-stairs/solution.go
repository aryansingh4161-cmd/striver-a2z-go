package climbingstairs

// Recursive: f(n) = f(n-1) + f(n-2). O(2^n).
func Recursive(n int) int {
	if n <= 1 {
		return 1
	}
	return Recursive(n-1) + Recursive(n-2)
}

// Memo: top-down with a table. O(n) time, O(n) space.
func Memo(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	var f func(int) int
	f = func(i int) int {
		if i <= 1 {
			return 1
		}
		if dp[i] != -1 {
			return dp[i]
		}
		dp[i] = f(i-1) + f(i-2)
		return dp[i]
	}
	return f(n)
}

// Tab: bottom-up table. O(n) time, O(n) space.
func Tab(n int) int {
	dp := make([]int, max(n+1, 2))
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Optimal: only the last two values. O(n) time, O(1) space.
func Optimal(n int) int {
	prev2, prev1 := 1, 1
	for i := 2; i <= n; i++ {
		prev2, prev1 = prev1, prev1+prev2
	}
	return prev1
}
