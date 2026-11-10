package unboundedknapsack

// Recursive: skip item i, or take it and stay on i. Exponential.
func Recursive(wt, val []int, W int) int {
	var f func(i, w int) int
	f = func(i, w int) int {
		if i == 0 {
			return w / wt[0] * val[0]
		}
		best := f(i-1, w)
		if wt[i] <= w {
			best = max(best, val[i]+f(i, w-wt[i]))
		}
		return best
	}
	return f(len(wt)-1, W)
}

// Memo caches f(i, w). O(n·W).
func Memo(wt, val []int, W int) int {
	dp := make([][]int, len(wt))
	for i := range dp {
		dp[i] = make([]int, W+1)
		for w := range dp[i] {
			dp[i][w] = -1
		}
	}
	var f func(i, w int) int
	f = func(i, w int) int {
		if i == 0 {
			return w / wt[0] * val[0]
		}
		if dp[i][w] != -1 {
			return dp[i][w]
		}
		best := f(i-1, w)
		if wt[i] <= w {
			best = max(best, val[i]+f(i, w-wt[i]))
		}
		dp[i][w] = best
		return best
	}
	return f(len(wt)-1, W)
}

// Optimal: one row, w going upward so items can repeat. O(W) space.
func Optimal(wt, val []int, W int) int {
	dp := make([]int, W+1)
	for i := range wt {
		for w := wt[i]; w <= W; w++ {
			dp[w] = max(dp[w], val[i]+dp[w-wt[i]])
		}
	}
	return dp[W]
}
