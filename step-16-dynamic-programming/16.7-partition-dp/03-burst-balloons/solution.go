package burstballoons

func pad(nums []int) []int {
	a := append([]int{1}, nums...)
	return append(a, 1)
}

// Recursive picks the LAST balloon burst in every range. Exponential.
func Recursive(nums []int) int {
	a := pad(nums)
	var f func(i, j int) int
	f = func(i, j int) int {
		if i > j {
			return 0
		}
		best := 0
		for k := i; k <= j; k++ {
			best = max(best, a[i-1]*a[k]*a[j+1]+f(i, k-1)+f(k+1, j))
		}
		return best
	}
	return f(1, len(nums))
}

// Tab is the bottom-up version. O(n³).
func Tab(nums []int) int {
	a := pad(nums)
	n := len(nums)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for i := n; i >= 1; i-- {
		for j := i; j <= n; j++ {
			best := 0
			for k := i; k <= j; k++ {
				best = max(best, a[i-1]*a[k]*a[j+1]+dp[i][k-1]+dp[k+1][j])
			}
			dp[i][j] = best
		}
	}
	return dp[1][n]
}
