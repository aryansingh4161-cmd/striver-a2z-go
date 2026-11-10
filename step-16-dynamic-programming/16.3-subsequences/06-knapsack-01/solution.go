package knapsack01

// Recursive: take or skip each item. O(2^n).
func Recursive(wt, val []int, W int) int {
	var f func(i, w int) int
	f = func(i, w int) int {
		if i == 0 {
			if wt[0] <= w {
				return val[0]
			}
			return 0
		}
		best := f(i-1, w)
		if wt[i] <= w {
			best = max(best, val[i]+f(i-1, w-wt[i]))
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
			if wt[0] <= w {
				return val[0]
			}
			return 0
		}
		if dp[i][w] != -1 {
			return dp[i][w]
		}
		best := f(i-1, w)
		if wt[i] <= w {
			best = max(best, val[i]+f(i-1, w-wt[i]))
		}
		dp[i][w] = best
		return best
	}
	return f(len(wt)-1, W)
}

// Tab is the full table. O(n·W) time and space.
func Tab(wt, val []int, W int) int {
	n := len(wt)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for w := wt[0]; w <= W; w++ {
		dp[0][w] = val[0]
	}
	for i := 1; i < n; i++ {
		for w := 0; w <= W; w++ {
			dp[i][w] = dp[i-1][w]
			if wt[i] <= w {
				dp[i][w] = max(dp[i][w], val[i]+dp[i-1][w-wt[i]])
			}
		}
	}
	return dp[n-1][W]
}

// Optimal uses one row with w going downward. O(W) space.
func Optimal(wt, val []int, W int) int {
	dp := make([]int, W+1)
	for i := range wt {
		for w := W; w >= wt[i]; w-- {
			dp[w] = max(dp[w], val[i]+dp[w-wt[i]])
		}
	}
	return dp[W]
}
