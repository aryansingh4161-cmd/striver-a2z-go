package matrixchainmultiplication

const inf = 1 << 60

// Recursive tries every split. Exponential.
func Recursive(arr []int) int {
	var f func(i, j int) int
	f = func(i, j int) int {
		if i == j {
			return 0
		}
		best := inf
		for k := i; k < j; k++ {
			best = min(best, f(i, k)+f(k+1, j)+arr[i-1]*arr[k]*arr[j])
		}
		return best
	}
	return f(1, len(arr)-1)
}

// Memo caches f(i, j). O(n³).
func Memo(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(i, j int) int
	f = func(i, j int) int {
		if i == j {
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		best := inf
		for k := i; k < j; k++ {
			best = min(best, f(i, k)+f(k+1, j)+arr[i-1]*arr[k]*arr[j])
		}
		dp[i][j] = best
		return best
	}
	return f(1, n-1)
}

// Tab fills shorter ranges first (i downward, j upward). O(n³).
func Tab(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j < n; j++ {
			best := inf
			for k := i; k < j; k++ {
				best = min(best, dp[i][k]+dp[k+1][j]+arr[i-1]*arr[k]*arr[j])
			}
			dp[i][j] = best
		}
	}
	return dp[1][n-1]
}
