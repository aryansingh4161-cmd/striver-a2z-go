package partitionarraymaxsum

// Recursive tries every first-block length. O(k^n).
func Recursive(arr []int, k int) int {
	n := len(arr)
	var f func(i int) int
	f = func(i int) int {
		if i == n {
			return 0
		}
		best, mx := 0, 0
		for l := 1; l <= k && i+l <= n; l++ {
			mx = max(mx, arr[i+l-1])
			best = max(best, mx*l+f(i+l))
		}
		return best
	}
	return f(0)
}

// Tab fills f(i) from the end. O(n·k) time, O(n) space.
func Tab(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		mx := 0
		for l := 1; l <= k && i+l <= n; l++ {
			mx = max(mx, arr[i+l-1])
			dp[i] = max(dp[i], mx*l+dp[i+l])
		}
	}
	return dp[0]
}
